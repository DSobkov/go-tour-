package main

import (
	"fmt"
	"html/template"
	"net/http"

)

func main() {
	fmt.Println("Listening on port :9949")
	http.ListenAndServe(":9949", nil)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/users", usersPage)
}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	// user := User{"Danylo", "Sobkov"}
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

func usersPage(w http.ResponseWriter, r *http.Request)  {
	users := []User{User{"Danylo", "Sobkov"}, User{"Vlad", "Orlov"}}
	// js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("users.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, users); err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
