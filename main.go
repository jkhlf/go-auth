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
