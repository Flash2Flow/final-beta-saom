package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Cardinal struct {
	Status string
}


func api_page(page http.ResponseWriter, req *http.Request) {
	page.Header().Set("Access-Control-Allow-Origin", "*")
	page.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Println("New request api")

	title := req.FormValue("title")
	login := req.FormValue("login")
	email := req.FormValue("email")
	password := req.FormValue("password")
	jwt := req.FormValue("jwt")

	if title == "auth" {
		log.Println("new req on auth page")
		if login == "" {
			mss := Cardinal{"Login nil"}
			js, err := json.Marshal(mss)
			if err != nil {
				http.Error(page, err.Error(), http.StatusInternalServerError)
				return
			}
			page.Header().Set("Content-Type", "application/json")
			page.Write(js)
			log.Println("Login nil")
		} else {
			if password == "" {
				mss := Cardinal{"Pass nil"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("Pass nil")
			} else {
				if jwt == ""{
					mss := Cardinal{"jwt nil"}
					js, err := json.Marshal(mss)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}
					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Println("jwt nil")
				}else{
					//get jwt
					JwtFindFile := fmt.Sprintf("server/jwt_auth.json")
					dat, err := ioutil.ReadFile(JwtFindFile)
					if err != nil {
						log.Println("Err 105 line")

					}
					JwtFull := JwtStruct{}
					err = json.Unmarshal(dat, &JwtFull)
					if err != nil {
						log.Println("line 111")
					}
					JwtNewType, err := strconv.Atoi(jwt)

					if int64(JwtNewType) == JwtFull.Token {
						err := Auth(login, password)
						if err != nil{
							log.Println(err.Error())
							msg := strings.HasPrefix(err.Error(), "ok, auth")
							log.Println(msg)
							if msg == true{
								mss := Cardinal{err.Error()}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
							}

							if err.Error() == "no, bad pass"{
								mss := Cardinal{"no, bad pass"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
								log.Println("no, bad pass")
							}

							if err.Error() == "no one user with this login"{
								mss := Cardinal{"no one user with this login"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
								log.Println("no one user with this login")
							}
						}
					}else{
						mss := Cardinal{"bad jwt"}
						js, err := json.Marshal(mss)
						if err != nil {
							http.Error(page, err.Error(), http.StatusInternalServerError)
							return
						}
						page.Header().Set("Content-Type", "application/json")
						page.Write(js)
						log.Println("bad jwt")
					}

					//get user
					// if err send bad req
					// if err == nil check pass
					// url pass == sever.user pass
					//if ok, send good req
					//else send bad req
				}
			}
		}
	}

	if title == "reg" {
		log.Println("new req on reg page")

		if login == "" {
			mss := Cardinal{"nil login"}
			js, err := json.Marshal(mss)
			if err != nil {
				http.Error(page, err.Error(), http.StatusInternalServerError)
				return
			}
			page.Header().Set("Content-Type", "application/json")
			page.Write(js)
			log.Println("nil login")
		} else {
			if password == "" {
				mss := Cardinal{"nil pass"}
				js, err := json.Marshal(mss)
				if err != nil {
					http.Error(page, err.Error(), http.StatusInternalServerError)
					return
				}
				page.Header().Set("Content-Type", "application/json")
				page.Write(js)
				log.Println("nil pass")
			} else {
				if email == "" {
					mss := Cardinal{"nil email"}
					js, err := json.Marshal(mss)
					if err != nil {
						http.Error(page, err.Error(), http.StatusInternalServerError)
						return
					}
					page.Header().Set("Content-Type", "application/json")
					page.Write(js)
					log.Println("nil email")
				} else {
					if jwt == "" {
						mss := Cardinal{"nil jwt"}
						js, err := json.Marshal(mss)
						if err != nil {
							http.Error(page, err.Error(), http.StatusInternalServerError)
							return
						}
						page.Header().Set("Content-Type", "application/json")
						page.Write(js)
						log.Println("nil jwt")
					} else {
						//server check
						JwtFindFile := fmt.Sprintf("server/jwt.json")
						dat, err := ioutil.ReadFile(JwtFindFile)
						if err != nil {
							log.Println("Err 105 line")

						}
						JwtFull := JwtStruct{}
						err = json.Unmarshal(dat, &JwtFull)
						if err != nil {
							log.Println("line 111")
						}
						JwtNewType, err := strconv.Atoi(jwt)

						if int64(JwtNewType) == JwtFull.Token {
							err := RegisterFindEmail(email)
							if err != nil {
								mss := Cardinal{"not new email"}
								js, err := json.Marshal(mss)
								if err != nil {
									http.Error(page, err.Error(), http.StatusInternalServerError)
									return
								}
								page.Header().Set("Content-Type", "application/json")
								page.Write(js)
								log.Println("not new email")
							} else {
								//check user
								err := RegisterFindUser(login)
								if err != nil {
									mss := Cardinal{"no new user"}
									js, err := json.Marshal(mss)
									if err != nil {
										http.Error(page, err.Error(), http.StatusInternalServerError)
										return
									}
									page.Header().Set("Content-Type", "application/json")
									page.Write(js)
									log.Println("no new user")
								} else {
									//register user

									err := Register(login, email, password)
									if err != nil {
										log.Println(err)
									}

									mss := Cardinal{"registration true"}
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
						} else {
							//bad jwt
							mss := Cardinal{"bad jwt"}
							js, err := json.Marshal(mss)
							if err != nil {
								http.Error(page, err.Error(), http.StatusInternalServerError)
								return
							}
							page.Header().Set("Content-Type", "application/json")
							page.Write(js)
							log.Println("bad jwt")
						}
					}
				}
			}
		}
	}

	if title == "update_key"{
		if login == ""{
			log.Println("Update user key cant be , because login nil")
		}else{
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
					return
				}

			UpdateAccess(UserFull.Login, UserFull.Password, UserFull.Email, UserFull.UserKey, UserFull.Developer, UserFull.Ban, UserFull.Group, UserFull.Undesirable)
		}else{
			log.Println("Haven't user")
		}

		}
	}


		if title == "" {
			log.Println("empty title")
		}

}