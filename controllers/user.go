package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dragosslujeru/go-rest-api/model"
)

var defaultUser = model.User{
	FirstName: "John",
	LastName:  "Smith",
	Email:     "john.smith@email.com"}

var users = []model.User{
	{FirstName: "John", LastName: "Doe", Email: "email1@email.com"},
	{FirstName: "Jane", LastName: "Doe", Email: "email2@email.com"},
	{FirstName: "Adam", LastName: "Smith", Email: "email3@email.com"},
}

//GetUser endpoint handler
var GetUser GetOnlyHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, defaultUser)
}

//GetAllUsers will return a list of all users
var GetAllUsers GetOnlyHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, struct {
		Users []model.User `json:"users"`
	}{users})
}

// PostUser will save a new user to the list
var PostUser PostOnlyHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u model.User
	decoder.Decode(&u)
	fmt.Printf("Created user %v\n", u)
	users = append(users, u)
}
