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



func JwtCreate() error{

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

	rand.Seed(time.Now().UnixNano())
	m := rand.Int63()
	log.Println(m)

	c := JwtStruct{
		Token: m,
		Date: time.Now(),
	}
	dat3, err := json.Marshal(c)
	if err != nil {
		return err
		log.Println("err50")
	}
	JwtRead := fmt.Sprintf("server/jwt_auth.json")
	err = ioutil.WriteFile(JwtRead, dat3, 0644)
	if err != nil {
		return err
		log.Println("err56")
	}
	log.Println("JWT update")

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
			errnew := fmt.Sprintf("ok, auth %d", UserFull.UserKey)
			return errors.New(errnew )
		}else{
			return errors.New("no, bad pass")
		}
	}else{
		return errors.New("no one user with this login")
	}

	return nil
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


	//update jwt

	rand.Seed(time.Now().UnixNano())
	m := rand.Int63()
	log.Println(m)

	c := JwtStruct{
		Token: m,
		Date: time.Now(),
	}
	dat3, err := json.Marshal(c)
	if err != nil {
		return err
		log.Println("err50")
	}
	JwtRead := fmt.Sprintf("server/jwt.json")
	err = ioutil.WriteFile(JwtRead, dat3, 0644)
	if err != nil {
		return err
		log.Println("err56")
	}
	log.Println("JWT update")
	return nil
}

func UpdateAccess(login string, password string, email string, UserKey int, dev int, ban int, group int, undesirable int) error{
	users := fmt.Sprintf("users/user_%s.json", login)
	f, err := os.Create(users)
	if err != nil {
		panic(err)
	}
	f.Close()
	u := UserFull{
		Login: login,
		Password: password,
		Email: email,
		Developer: dev,
		Ban: ban,
		Group: group,
		Undesirable: undesirable,
		UserKey: UserKey,
	}
	dat, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(users, dat, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}