package main

import (
	"context"
	"fmt"
	"github.com/go-session/session"
	"log"
	"net/http"
)

func main()  {
	//var server settings
		var SettingsPort = "server started with port 8080\r\n"
		var SettingsVersion = "version beta 0.8.356"
	//end var server settings

	//log server settings
		log.Println(SettingsPort + SettingsVersion)
	//end log

	//start server
		server()
	//end start server
}

func server()  {
	//server func, handle all pages

		//home page = auth and reg func
			http.HandleFunc("/", home)
		//end home page

		//home active page
			http.HandleFunc("/home/", homeActive)
		//end home active page

		//func api, handle all request for auth and reg
			http.HandleFunc("/api", ApiPage)
		//end func api

		//func auth, get on home page, and redirect on home active or home
			http.HandleFunc("/auth", auth)
		//end func auth

		//func delete = delete session
			http.HandleFunc("/delete/", delete)
		//end func delete

		//func handle static files ( html, css, image, js... )
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
		//end handle static files

		//func server start
			http.ListenAndServe(":8080", nil)
		//end func server start

	//end server func
}

func delete(page http.ResponseWriter, req *http.Request){
	//var for slow redirect, in future change on js redirect
		var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=https://92.255.104.121/\" /> </head>"
	//end var redirect

	//start session
		store, err := session.Start(context.Background(), page, req)
		if err != nil {
			fmt.Fprint(page, err)
			return
		}
	//end session

	//get active session
		active, ok := store.Get("active_login")
	//end get session active

	//check session active
		if ok {
			//delete active session
				session.Destroy(context.Background(), page, req)
				store.Delete("active_login")
			//end delete active session

			//log about delete session
			Active := fmt.Sprintf("User del.s: %v", active)
			log.Println(Active)
			//end log delete

			//redirect
				fmt.Fprintf(page, Redirect)
			//end redirect
		}else{
			//redirect
				fmt.Fprintf(page, Redirect)
			//end redirect
		}
	//end check session active
}