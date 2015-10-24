package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"State/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

// GetUser retrieves an individual user resource
func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get user
	user := models.User{
	    Name:   "John Doe",
	    Email: "johndoe@email.com",
	    Id:     p.ByName("id"),
	}

	userjson, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", userjson)
}

// CreateUser creates a new user resource
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	// create user


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "Created User")
}

// RemoveUser removes an existing user resource
func (userController UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	// delete user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "Deleted User")
}