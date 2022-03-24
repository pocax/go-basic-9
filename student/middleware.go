package main

import (
	"net/http"
)

const USERNAME = "root"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Unauthorized"))
		return false
	}
	isValid := (username == USERNAME && password == PASSWORD)
	if !isValid {
		w.Write([]byte("Unauthorized"))
		return false
	}
	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Method not allowed"))
		return false
	}
	return true
}
