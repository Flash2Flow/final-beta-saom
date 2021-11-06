package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-session/session"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func auth(page http.ResponseWriter, req *http.Request){
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://45.128.207.175:8010/home/\" /> </head>"
	_, ok := store.Get("active_login")



	//if have active session, use active pages
	if ok {

		//redirect on active page
		fmt.Fprintf(page, Redirect)

	}

	login := req.FormValue("login")
	access := req.FormValue("access")

	if login == ""{
		log.Println("empty login")
	}else{
		if access == ""{
			log.Println("empty access")
		}else{
			//get access token
			JwtFindFile := fmt.Sprintf("users/user_%s.json", login)
			dat, err := ioutil.ReadFile(JwtFindFile)
			if err != nil {
				log.Println("Err 105 line")

			}
			User := UserFull{}
			err = json.Unmarshal(dat, &User)
			if err != nil {
				log.Println("line 111")
			}
			// check oba
			UserKey := strconv.Itoa(User.UserKey)
			if UserKey == access {
				store.Set("active_login", login)
				err = store.Save()
				if err != nil {
					fmt.Fprint(page, err)
					return
				}
				http.Redirect(page, req, "/home/", 302)
				auth := fmt.Sprintf("User auth: %s", login)
				log.Println(auth)
			}else{
				//redirect on home
				log.Println("bad access")
				http.Redirect(page, req, "/", 302)
			}
			//create session and redirect
		}
	}
}