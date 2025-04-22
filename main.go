package main

import (
	"net/http"
)

type Login struct {
	HashedPassword string
	SessionToken string
	CSRFToken string
}

var users = map[string]Login {

}

func main() {
    http.HandleFunc("/register", register)
		http.HandleFunc("/login", login)
		http.HandleFunc("/logout", logout)
    http.HandleFunc("/protected", protected)
		http.ListenAndServe(":8080", nil)
}

func register (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
}
}

func login (w http.ResponseWriter, r *http.Request) {}

func logout (w http.ResponseWriter, r *http.Request) {}


func protected (w http.ResponseWriter, r *http.Request) {}



