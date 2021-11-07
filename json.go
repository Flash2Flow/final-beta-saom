package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)


type JwtStruct struct {
	Token int64
	Date time.Time
}
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

func CreateJwt(value int64) error{
	JwtPath := fmt.Sprintf("server/jwt/jwt_%d.json", value)
	j := JwtStruct{
		Token: value,
		Date: time.Now(),
	}

	dat, err := json.Marshal(j)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(JwtPath, dat, 0644)
	if err != nil {
		return err
	}

	auth := fmt.Sprintf("Jwt create: %d", value)
	log.Println(auth)
	return nil
}

func RegisterFindUser(login string) error{
	users := fmt.Sprintf("users/user_%s.json", login)
	files, _ := filepath.Glob(users)
	if files != nil {
		return errors.New("no new user")
	}else{
		return nil
	}
}

func RegisterFindEmail(email string) error{
	emails := fmt.Sprintf("emails/email_%s.json", email)
	files, _ := filepath.Glob(emails)
	if files != nil {
		return errors.New("no new email")
	}else{
		return nil
	}
}

func Auth(login string, password string) error {

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
			return err
		}
		Md5Password := GetMD5Hash(password)
		if UserFull.Password == Md5Password {
			return errors.New("ok, can auth")
		}else{
			return errors.New("no, bad pass")
		}
	}else{
		return errors.New("no one user with this login")
	}

	return nil
}

func AuthFindLogin(login string, password string)  error {
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
			return err
		}
		Md5Password := GetMD5Hash(password)
		if UserFull.Password == Md5Password{
			return errors.New("ok, can auth")
		}else{
			return errors.New("bad password")
		}

	}else{
		return errors.New("have not user")
	}
}

func Register(login string, email string, password string) error{

	//create user json
	users := fmt.Sprintf("users/user_%s.json", login)
	f, err := os.Create(users)
	if err != nil {
		panic(err)
	}
	f.Close()


	Md5UserKey := rand.Intn(9999999999)

	Md5Password := GetMD5Hash(password)
	//reg
	u := UserFull{
		Login: login,
		Password: Md5Password,
		Email: email,
		Developer: 0,
		Ban: 0,
		Group: 0,
		Undesirable: 0,
		UserKey: Md5UserKey,
	}
	dat, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(users, dat, 0644)
	if err != nil {
		return err
	}


	//create email

	emails := fmt.Sprintf("emails/email_%s.json", email)
	f2, err2 := os.Create(emails)
	if err != nil {
		panic(err2)
	}
	f2.Close()


	//create email json
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63()

	//reg
	e := EmailFull{
		Email: email,
		Date: time.Now(),
		Key: n,
	}
	dat2, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(emails, dat2, 0644)
	if err != nil {
		return err
	}

	log.Println("Email create: " + email)
	log.Println("User create: " + login)

	return nil
}

func GetMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}