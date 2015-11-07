package controllers

import (
	"encoding/json"
	. "fmt"
	"net/http"

	"github.com/4nthem/State/models"
	"github.com/go-martini/martini"
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

func NewUserController() *UserController {
	return &UserController{}
}

///JUST HERE FOR NOW

func TestMiddleware1(w http.ResponseWriter, r *http.Request, p martini.Params, c martini.Context) {
	p["IsAuthorized"] = "true"
	c.Next() //Interesting. Just experimenting here. Next() will block till all other middleware have finished
	Printf("request1 finished!!")

}

func TestMiddleware2(w http.ResponseWriter, r *http.Request, p martini.Params, c martini.Context) {
	Printf("request2 finisehd!!")
	Fprintf(w, "AUTHORIZED ONLY: %s", p)
}

////////////////////////////

func (UserController UserController) Home(w http.ResponseWriter, r *http.Request, p martini.Params) {
	Fprintf(w, "params: %s", p)
}

func (userController UserController) GetUser(w http.ResponseWriter, r *http.Request, p martini.Params) {
	Fprintf(w, "Got Specific User: %s", p)
}

// gets an individual user
func (userController UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, p martini.Params) {

	// get user
	user := models.User{
		// Id:     params.ByName("id"),
		Id:    bson.NewObjectId(),
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	userjson, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	Fprintf(w, "%s", userjson)
}

// creates a new user
func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, p martini.Params) {

	// create user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	Fprintf(w, "Created User")
}

// removes an existing user
func (userController UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p martini.Params) {

	// delete user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	Fprintf(w, "Deleted User")
}
