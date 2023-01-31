package mockinggo

import (
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := &User{Name: "galih agus setyawan", Address: "trenggalek"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]User, 2)
	users = append(users, User{Name: "galih agus", Address: "trenggalek"}, User{Name: "agus", Address: "trenggalek"})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
