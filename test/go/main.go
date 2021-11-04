package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main(){
	log.Println("server started with port 1010")
	http.HandleFunc("/", home)
	http.ListenAndServe(":1010", nil)
}

func home(page http.ResponseWriter, req *http.Request)  {

	temp, err := template.ParseFiles("temp/home.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)
}