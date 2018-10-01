package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	if email == "" {
		http.Error(w, "The email can not be a null parameter", http.StatusBadRequest)
		return
	}

	u := FetchUserByEmail(email)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(u)
	w.Write([]byte(payload))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var u User
	err := decoder.Decode(&u)
	if err != nil {
		log.Println(err)
	}

	CreateNewUser(u)

	w.Header().Set("Content-Type", "text")
	fmt.Fprint(w, u)
}
