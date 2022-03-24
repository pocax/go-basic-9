package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OutputJson(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJson(w, SelectStudent(id))
		return
	}
	OutputJson(w, GetStudent())
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8181"

	fmt.Println("Server is running at port 8181")
	server.ListenAndServe()
}
