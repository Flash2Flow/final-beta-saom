package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"fmt"
	"github.com/go-session/session"

	"log"
	"net/http"
)

func auth(page http.ResponseWriter, req *http.Request) {
	//var for slow redirect, in future change on js redirect
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=https://92.255.104.121/\" /> </head>"
	//end var redirect

	page.Header().Set("Access-Control-Allow-Origin", "*")
	page.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	var RedirectHome = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=https://92.255.104.121/home/\" /> </head>"
	_, ok := store.Get("active_login")

	//if have active session, use active pages
	if ok {
		//redirect on active page
		fmt.Fprintf(page, RedirectHome)

	}

	login := req.FormValue("login")
	password := req.FormValue("password")


	_, ok2 := store.Get("jwt")
	if ok2 {

		users := fmt.Sprintf("users/user_%s.json", login)
		files, _ := filepath.Glob(users)
		if files != nil {
			dat, err := ioutil.ReadFile(users)
			if err != nil {
				//err unknown
			}

			UserFull := UserFull{}

			err = json.Unmarshal(dat, &UserFull)
			if err != nil {
				return
			}
			Md5Password := GetMD5Hash(password)
			if UserFull.Password == Md5Password {
				switch login {
				case "":
					log.Println("Login nil")
					fmt.Fprintf(page, Redirect)

				case UserFull.Login:
					store.Set("active_login", login)
					err = store.Save()
					if err != nil {
						fmt.Fprint(page, err)
						return
					}

					logAuth := fmt.Sprintf("User logged: %s", login)
					log.Println(logAuth)
					fmt.Fprintf(page, RedirectHome)
				}
			}else{
				log.Println("Bad pass")
				fmt.Fprintf(page, Redirect)
			}
		}else{
			log.Println("Have not this user")
			fmt.Fprintf(page, Redirect)
		}

	}
}