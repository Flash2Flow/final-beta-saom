package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-session/session"
)

type JWTstruct struct {
	Token int64
	Date time.Time
}

func home(page http.ResponseWriter, req *http.Request)  {
	log.Println("new request on home page")
			//variable with bad html redirect
			var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://localhost:5050/home/\" /> </head>"
			//var Redirect = "<script>window.location.replace('/home/');</script>" на будущее, работает только с https, но работает ахуенно

	//start session
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}

	//parse default html files on page
	temp, err := template.ParseFiles("temp/html/home.html", "temp/html/pre.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Int63()
	log.Println(n)

	c := JWTstruct{
		Token: n,
		Date: time.Now(),
	}
	dat, err := json.Marshal(c)
	if err != nil {
		return
		log.Println("err50")
	}
	jwt_read := fmt.Sprintf("server/jwt.json")
	err = ioutil.WriteFile(jwt_read, dat, 0644)
	if err != nil {
		return
		log.Println("err56")
	}

	ne := uint64(n)
	//exec html var on page
	temp.ExecuteTemplate(page, "home_page", struct{Token uint64}{Token: ne})

	//get active session, if empty, use default
	active, ok := store.Get("active_login")



	//if have active session, use active pages
	if ok {

		//log active user on bad for him page
		Active := fmt.Sprintf("HOME PAGE User logged: %v", active)
		log.Println(Active)

		//redirect on active page
		fmt.Fprintf(page, Redirect)

	}

}

func HomeActive(page http.ResponseWriter, req *http.Request) {
	log.Println("new request on home active page")
	//variable with bad html redirect
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://localhost:5050/\" /> </head>"
	//var Redirect = "<script>window.location.replace('/home/');</script>" на будущее, работает только с https, но работает ахуенно

	//start session
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}

	//parse default html files on page
	temp, err := template.ParseFiles("temp/html/home_a.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	//exec html var on page
	temp.ExecuteTemplate(page, "home_page_active", nil)

	//get active session, if empty, use default
	active, ok := store.Get("active_login")

	//if have active session, use active pages
	if ok {

		//log active user on bad for him page
		Active := fmt.Sprintf("HOME PAGE User unlogged: %v", active)
		log.Println(Active)

		temp.ExecuteTemplate(page, "home_page_active", active)

	}else{
		//redirect on default page
		fmt.Fprintf(page, Redirect)
	}

}