package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/4nthem/State/models"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	UserController struct {
		// session *mgo.Session
	}
)

// //  provides a reference to a UserController with provided mongo session
// func NewUserController(session *mgo.Session) *UserController {
// 	return &UserController{session}
// }

func NewUserController() *UserController{
	return &UserController{}
}

// gets an individual user 
func (userController UserController) GetUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// get user
	user := models.User{
		// Id:     params.ByName("id"),
		Id: 	bson.NewObjectId(),
	    Name:   "John Doe",
	    Email:  "johndoe@email.com",
	}

	userjson, _ := json.Marshal(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	fmt.Fprintf(writer, "Found User: %s", userjson)
}

// creates a new user 
func (userController UserController) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	// create user


	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	fmt.Fprintf(writer, "Created User")
}

// removes an existing user 
func (userController UserController) RemoveUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	// delete user

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	fmt.Fprintf(writer, "Deleted User")
}
