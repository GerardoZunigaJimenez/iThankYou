package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
