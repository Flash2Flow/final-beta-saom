package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Cardinal struct {
	Status string
}


func ApiPage(page http.ResponseWriter, req *http.Request) {
	page.Header().Set("Access-Control-Allow-Origin", "*")
	page.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("New request api")

	title := req.FormValue("title")
	login := req.FormValue("login")
	email := req.FormValue("email")
	password := req.FormValue("password")


	switch title {
	case "":
		log.Println("empty title")
		return

	case "reg":
		log.Println("new req on reg")
			if email == ""{
				log.Println("Empty email")
				return
			}

			if login == ""{
				log.Println("Empty login")
				return
			}

			if password == ""{
				log.Println("Empty password")
				return
			}

		err := RegisterFindUser(login)
		if err != nil{
			switch err.Error() {
			case "no new user":
				//send req
				mss := Cardinal{"no new user"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("no new user")

			case "true":

				err := RegisterFindEmail(email)
				if err != nil{
					switch err.Error() {
					case "no new email":
						//send req
						mss := Cardinal{"no new email"}
						js, err := json.Marshal(mss)
						if err != nil {
							http.Error(page, err.Error(), http.StatusInternalServerError)
							return
						}
						page.Header().Set("Content-Type", "application/json")
						page.Write(js)
						log.Println("no new email")

					case "ok":
						err := Register(login, email, password)
						if err != nil {
							log.Println(err)
						}

						mss := Cardinal{"ok, reg"}
						js, err := json.Marshal(mss)
						if err != nil {
							http.Error(page, err.Error(), http.StatusInternalServerError)
							return
						}
						page.Header().Set("Content-Type", "application/json")
						page.Write(js)
						log.Println("registration true")
					}
				}
			}
		}

	case "auth":
		log.Println("new req on auth")
		if login == ""{
			log.Println("empty login")
		}

		err := AuthFindLogin(login, password)
		if err != nil {
			switch err.Error() {
			case "have not user":
				mss := Cardinal{"have not user"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("have not user")

			case "ok, can auth":
				mss := Cardinal{"ok, can auth"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("can auth")

			case "bad password":
				mss := Cardinal{"bad password"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("bad password")
			}


		}
	}


}