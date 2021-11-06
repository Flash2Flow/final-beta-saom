package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-session/session"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)
func home(page http.ResponseWriter, req *http.Request)  {

	//parse default html files on page
	temp, err := template.ParseFiles("temp/html/home.html", "temp/html/pre.html")
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://localhost:1010/home/\" /> </head>"
	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	_, ok := store.Get("active_login")

	if ok {
		fmt.Fprintf(page, Redirect)
	}


	JwtFindFile := fmt.Sprintf("server/jwt.json")
	dat, err := ioutil.ReadFile(JwtFindFile)
	if err != nil {
		log.Println("Err 105 line")

	}
	JwtFull := JwtStruct{}
	err = json.Unmarshal(dat, &JwtFull)
	if err != nil {
		log.Println("line 111")
	}

	ne := uint64(JwtFull.Token)

	JwtFindFileAuth := fmt.Sprintf("server/jwt_auth.json")
	dat5, err := ioutil.ReadFile(JwtFindFileAuth)
	if err != nil {
		log.Println("Err 105 line")

	}
	err = json.Unmarshal(dat5, &JwtFull)
	if err != nil {
		log.Println("line 111")
	}

	nea := uint64(JwtFull.Token)

	//exec html var on page
	temp.ExecuteTemplate(page, "home_page",  struct{Token uint64; TokenAuth uint64}{Token: ne, TokenAuth: nea})



}

func homeActive(page http.ResponseWriter, req *http.Request)  {

	//parse default html files on page
	temp, err := template.ParseFiles("temp/html/home_a.html")
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://localhost:1010/\" /> </head>"
	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	active, ok := store.Get("active_login")


	if ok {
		temp.ExecuteTemplate(page, "home_page_active",  active)
	}else{
		fmt.Fprintf(page, Redirect)
	}



}