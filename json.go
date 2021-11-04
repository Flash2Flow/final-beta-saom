package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func jwt(time_now time.Time, random int64){
	rand.Seed(random)
	jwt := rand.Int63()
	log.Println(jwt)

	c := JWTstruct{
		Token: random,
		Date: time_now,
	}
	dat, err := json.Marshal(c)
	if err != nil {
		return
		log.Println("err50")
	}
	jwt_read := fmt.Sprintf("../server/jwt.json")
	err = ioutil.WriteFile(jwt_read, dat, 0644)
	if err != nil {
		return
		log.Println("err56")
	}

}