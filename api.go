package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

type Cardinal struct {
	Status string
	Code   string
}

const token = "3c8b160988b20198af262dc10bd2812fbac1a402f06d73270162a40866e39b769caac13177e2fd78c7183"

var vk = api.NewVK(token)

func api_page(page http.ResponseWriter, req *http.Request){

	page.Header().Set("Content-Type", "text/html; charset=utf-8")
	page.Header().Set("Access-Control-Allow-Origin", "*")

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	//longpoll
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}



	log.Println("New request api")

	title := req.FormValue("title")
	login := req.FormValue("login")
	email_url := req.FormValue("email")
	password := req.FormValue("password")
	jwt := req.FormValue("jwt")



	//registration

		//get url values

		//check login/pass/email

		//check jwt token for login

		//if jwt success => rest reg.go

		//else

		//redirect bad page

	if title == "reg"	{
		if login == ""{
			mss := Cardinal{"Login nil", "201"}
			js, err := json.Marshal(mss)
			if err != nil {
				http.Error(page, err.Error(), http.StatusInternalServerError)
				return
			}
			page.Header().Set("Content-Type", "application/json")
			page.Write(js)
			log.Print("Login nil: 201")
		}else{
			if password == ""{
				mss := Cardinal{"Password nil", "202"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Print("Password nil: 202")
			}else{
				if email_url == "" {
					mss := Cardinal{"Email nil", "203"}
					js, err := json.Marshal(mss)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}
					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Print("Email nil: 203")
				}else{
					if jwt == ""{
						mss := Cardinal{"JWT nil", "204"}
						js, err := json.Marshal(mss)
						if err != nil {
							http.Error(page, err.Error(), http.StatusInternalServerError)
							return
						}
						page.Header().Set("Content-Type", "application/json")
						page.Write(js)
						log.Print("JWT nil: 204")
					}else{
						users := fmt.Sprintf("server/jwt.json")
						dat, err := ioutil.ReadFile(users)
						if err != nil {
							//nil jwt
							log.Println("Err 115 line")
							log.Println(err)
						}

						jwt_full := JWTstruct{}

						err = json.Unmarshal(dat, &jwt_full)
						if err != nil {
							return
						}
						jwt_suka, err := strconv.Atoi(jwt)
						if jwt_full.Token == int64(jwt_suka) {
							users := fmt.Sprintf("users/user_%s.json", login)
							files, _ := filepath.Glob(users)
							if files != nil {
								mss := Cardinal{"User already have", "100"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
								log.Print("User already have: 100")
							}else{
								email := fmt.Sprintf("/emails/email_%s.json", email_url)
								files, _ := filepath.Glob(email)
								if files != nil {
									mss := Cardinal{"Email already have", "102"}
									js, err := json.Marshal(mss)
									if err != nil {
										http.Error(page, err.Error(), http.StatusInternalServerError)
										return
									}
									page.Header().Set("Content-Type", "application/json")
									page.Write(js)
									log.Print("Email already have: 102")
								}else{
									//create user
									users := fmt.Sprintf("users/user_%s.json", login)
									f, err := os.Create(users)
									if err != nil {
										panic(err)
									}
									f.Close()

									md5_userkey := rand.Intn(9999999999)

									md5_password := GetMD5Hash(password)
									//reg
									c := UserFull{
										Login: login,
										Password: md5_password,
										Email: email_url,
										Developer: 0,
										Ban: 0,
										Group: 0,
										Undesirable: 0,
										UserKey: md5_userkey,
									}
									dat, err := json.Marshal(c)
									if err != nil {
										return
									}
									users_read := fmt.Sprintf("users/user_%s.json", login)
									err = ioutil.WriteFile(users_read, dat, 0644)
									if err != nil {
										return
									}

									//create email

									email := fmt.Sprintf("emails/email_%s.json", email_url)
									f2, err2 := os.Create(email)
									if err != nil {
										panic(err2)
									}
									f2.Close()

									rand.Seed(time.Now().UnixNano())
									n := rand.Int63()

									//reg
									e := EmailFull{
										Email: email_url,
										Date: time.Now(),
										Key: n,
									}
									dat2, err2 := json.Marshal(e)
									if err2 != nil {
										return
									}
									email_read := fmt.Sprintf("emails/email_%s.json", email_url)
									err = ioutil.WriteFile(email_read, dat2, 0644)
									if err != nil {
										return
									}

									email_c := fmt.Sprintf("Email create: %s", email_url)
									log.Println(email_c)

									user_c := fmt.Sprintf("User create: %s", login)
									log.Println(user_c)
								}
							}
						}else{
							//jwt bad
							log.Println("bad jwt")
						}

					}
				}
			}
		}

	}


	if title == "auth" {
		if login == ""{
			mss := Cardinal{"Login nil", "201"}
			js, err := json.Marshal(mss)
			if err != nil {
				http.Error(page, err.Error(), http.StatusInternalServerError)
				return
			}
			page.Header().Set("Content-Type", "application/json")
			page.Write(js)
			log.Print("Login nil: 201")
		}else{
			if password == ""{
				mss := Cardinal{"Password nil", "202"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Print("Password nil: 202")
			}else{
				if jwt == "" {
					mss := Cardinal{"JWT nil", "204"}
					js, err := json.Marshal(mss)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}
					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Print("JWT nil: 204")
				}else{
					//auth
					jwt_find := fmt.Sprintf("server/jwt.json")
					dat, err := ioutil.ReadFile(jwt_find)
					if err != nil {
						log.Println("Err 115 line")
						log.Println(err)
					}

					jwt_full := JWTstruct{}

					err = json.Unmarshal(dat, &jwt_full)
					if err != nil {
						return
					}
					log.Println(jwt)
					jwt_suka, err := strconv.Atoi(jwt)


					if jwt_full.Token == int64(jwt_suka){
						users := fmt.Sprintf("users/user_%s.json", login)
						files, _ := filepath.Glob(users)
						if files != nil {
							dat, err := ioutil.ReadFile(users)
							if err != nil {
								//err unknown
							}

							users_full := UserFull{}

							err = json.Unmarshal(dat, &users_full)
							if err != nil {
								return
							}

							if users_full.Password == password {
								mss := Cardinal{"Access auth", "1"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
								log.Print("Access auth: 1")
							}else{
								mss := Cardinal{"Pass inc", "302"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									log.Println(err)
									return
								}

								page.Write(js)
								page.Header().Set("Content-Type", "application/json")
								log.Print("Pass inc: 302")
							}
						}else{
							mss := Cardinal{"Empty user", "101"}
							js, err := json.Marshal(mss)
							if err != nil {
								http.Error(page, err.Error(), http.StatusInternalServerError)
								return
							}
							page.Header().Set("Content-Type", "application/json")
							page.Write(js)
							log.Print("Empty user: 101")

							log.Println("Нет такого человека блять")
						}
					}else{
						log.Println("bad jwt")
					}

				}
			}
		}
	}

	if title == ""{
		dev := "380236476"
		log.Println("Trying connect to empty site")

		_, err := vk.MessagesSend(api.Params{
			"peer_id":   dev,
			"random_id": 0,
			"message":   "Trying connect to empty site",
		})
		if err != nil {
			log.Println("160 line err")
		}
	}

	//authorization

		//get url values

		//check login/pass

		//check jwt token for login

		//if jwt success => rest auth.go

		//else

		//redirect bad page

	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}


func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}