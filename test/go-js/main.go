package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseTest struct {
	Status		string
	Code string
}

func main(){
	log.Println("server started with port 2020")
	http.HandleFunc("/api", api)
	http.ListenAndServe(":2020", nil)
}

func api(page http.ResponseWriter, req *http.Request)  {

	page.Header().Set("Content-Type", "text/html; charset=utf-8")
	page.Header().Set("Access-Control-Allow-Origin", "*")

	login := req.FormValue("login")
	pass := req.FormValue("password")

	log.Println("suka" + login + pass)

	if login == "" {
		//err
		responsetest := ResponseTest{"ERR",  "1 Login empty"}

		js, err := json.Marshal(responsetest)
		if err != nil {
			http.Error(page, err.Error(), http.StatusInternalServerError)
			return
		}

		page.Header().Set("Content-Type", "application/json")
		page.Write(js)
		log.Println(js)
	}else{
		if login == "yes"{
			if pass == ""{
				//err
				responsetest := ResponseTest{"ERR",  "2 Pass empty"}

				js, err := json.Marshal(responsetest)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}

				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println(js)
			}else{
				if pass == "1234"{
					//yes
					responsetest := ResponseTest{"YES",  "1 You are gay"}

					js, err := json.Marshal(responsetest)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}

					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Println(js)
				}else{
					//err
					responsetest := ResponseTest{"ERR",  "4 Pass Invalid"}

					js, err := json.Marshal(responsetest)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}

					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Println(js)
				}
			}
		}else{
			//err
			responsetest := ResponseTest{"ERR",  "3 Login invalid"}

			js, err := json.Marshal(responsetest)
			if err != nil {
				http.Error(page, err.Error(), http.StatusInternalServerError)
				return
			}

			page.Header().Set("Content-Type", "application/json")
			page.Write(js)
			log.Println(js)
		}
	}
}