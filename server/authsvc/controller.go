package authsvc

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/pomodoro/db"
	"github.com/J-Obog/pomodoro/usersvc"
	"golang.org/x/crypto/bcrypt"
)

func LogUserIn(w http.ResponseWriter, r *http.Request) {

}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user usersvc.User


	if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
		w.WriteHeader(500)
		return
	}
	
	if hash, e := bcrypt.GenerateFromPassword([]byte(user.Password), 10); e != nil {
		user.Password = string(hash)
	}
	
	if e := db.DB.Create(&user).Error; e != nil {
		w.WriteHeader(500)
		return
	}

	if res, e := json.Marshal(user); e != nil {
		w.WriteHeader(500)
		return
	} else {
		w.Write(res)
	}

}

func LogUserOut(w http.ResponseWriter, r *http.Request) {

}