package main

import (
	"fmt"
	"html/template"
	"net/http"

)

func main() {
	fmt.Println("Listening on port :9999")
	http.ListenAndServe(":9999", nil)
	http.HandleFunc("/", mainPage)

}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	// user := User{"Danylo"}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
