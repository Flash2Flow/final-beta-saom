package main

import (
	"context"
	"fmt"
	"github.com/go-session/session"
	"log"
	"net/http"
	"time"
)

type UserFull struct {
	Login   		  string
	Email 			  string
	Password 		  string
	Developer   	  int
	Ban				  int
	Group       	  int
	Undesirable 	  int
	UserKey			  int
}

type EmailFull struct{
	Email string
	Date time.Time
	Key int64
}

func main(){
	log.Println("server started with port 1010")
	server()
}

func server(){
	http.HandleFunc("/", home)
	http.HandleFunc("/home/", HomeActive)
	http.HandleFunc("/api", api_page)
	http.HandleFunc("/delete/", delete)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":1010", nil)
}

func delete(page http.ResponseWriter, req *http.Request){
	var Redirect = "<head> <meta http-equiv=\"refresh\" content=\"1;URL=http://localhost:5050/home/\" /> </head>"

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	active, ok := store.Get("active_login")

	if ok {
		Active := fmt.Sprintf("User del.s: %v", active)
		log.Println(Active)
		session.Destroy(context.Background(), page, req)
		store.Delete("active_login")
		fmt.Fprintf(page, Redirect)
	}else{
		fmt.Fprintf(page, Redirect)
	}

}