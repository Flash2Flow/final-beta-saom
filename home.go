package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-session/session"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)
func home(page http.ResponseWriter, req *http.Request)  {
	//var for slow redirect, in future change on js redirect
		var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=https://92.255.104.121/home/\" /> </head>"
	//end var redirect

	//var for reload page
		var Reload = "<script type=\"text/javascript\">window.location.reload();</script>"
	//end var reload

	//headers for idiots CORS
		page.Header().Set("Access-Control-Allow-Origin", "*")
		page.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//end headers

	//parse default html files on page
		temp, err := template.ParseFiles("temp/html/home.html", "temp/html/pre.html")
		if err != nil {
			fmt.Fprintf(page, err.Error())
		}
	//end parse default html files

	//start session
		store, err := session.Start(context.Background(), page, req)
		if err != nil {
			fmt.Fprint(page, err)
			return
		}
	//end session start

	//get session active
		_, ok := store.Get("active_login")
	//end get session active

	//check session active
		if ok {
			fmt.Fprintf(page, Redirect)
		}
	//end check session active

	//get session jwt
		JwtSession, ok2 := store.Get("jwt")
	//end get jwt session

	//check jwt session
	if ok2 {
		//find jwt file
			users := fmt.Sprintf("server/jwt/jwt_%d.json", JwtSession)
			files, _ := filepath.Glob(users)
				if files == nil{
				//delete session jwt because him bad
					store.Delete("jwt")
				//end delete session

				//reload page
					fmt.Fprintf(page, Reload)
				//end reload page
				}
		//end find jwt file

		//get values in file
		dat, err := ioutil.ReadFile(users)
			if err != nil {
				//err unknown
			}

		JwtFull := JwtStruct{}

		err = json.Unmarshal(dat, &JwtFull)
			if err != nil {
				return
			}
		//end get values

		//check values in jwt file with jwt session

		if JwtFull.Token != JwtSession {

			//delete session jwt because him bad
				store.Delete("jwt")
			//end delete jwt session

			//delete jwt file because him bad
				os.Remove(users)
			//end delete jwt file

			//reload page
				fmt.Fprintf(page, Reload)
			//end reload page

		}

		//end check values

	}else{
		//create jwt file

			//create random jwt
				rand.Seed(time.Now().UnixNano())
				s := rand.Int63()
			//end create random jwt

			//find files
				users := fmt.Sprintf("server/jwt/jwt_%d.json", s)
				files, _ := filepath.Glob(users)
			//end find files

				//check file for have
					if files == nil{
						//create file
							f, err := os.Create(users)
							if err != nil {
								panic(err)
							}
							f.Close()

						err2 := CreateJwt(s)
						if err2 != nil{
							log.Println("err line 134")
						}

						//end create users
					}
				//end check file for have

		//end create jwt file

		//set jwt session

			store.Set("jwt", s)
			err = store.Save()
			if err != nil {
				fmt.Fprint(page, err)
				return
			}
		//end set jwt session
	}


	JwtNewType := fmt.Sprintf("%v", JwtSession)
	//exec html var on page
		temp.ExecuteTemplate(page, "home_page",  struct{Jwt string}{Jwt: JwtNewType})
	//end exec

	//struct{Token uint64; TokenAuth uint64}{Token: ne, TokenAuth: nea}
}

func homeActive(page http.ResponseWriter, req *http.Request)  {
	//var for slow redirect, in future change on js redirect
		var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=https://92.255.104.121/\" /> </head>"
	//end var redirect


	//headers for idiots CORS
		page.Header().Set("Access-Control-Allow-Origin", "*")
		page.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//end headers

	//parse default html files on page
		temp, err := template.ParseFiles("temp/html/home_a.html")
		if err != nil {
			fmt.Fprintf(page, err.Error())
		}
	//end parse

	//start session
		store, err := session.Start(context.Background(), page, req)
		if err != nil {
			fmt.Fprint(page, err)
			return
		}
	//end session

	//get active session
		active, ok := store.Get("active_login")
	//end get active session

	//check session active
		if ok {
			//true session, exec login
				temp.ExecuteTemplate(page, "home_page_active",  active)
			//end exec
		}else{
			//empty session, redirect
				fmt.Fprintf(page, Redirect)
			//end redirect
		}
	//end check session

	//get session jwt
		_, ok2 := store.Get("jwt")
	//end get session jwt

	//check session jwt
		if ok2 {
			//delete session jwt because him bad
			store.Delete("jwt")
			//end delete jwt session
		}
	//end check session

}