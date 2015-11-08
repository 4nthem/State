package controllers

import (
	"net/http"

	"github.com/4nthem/State/auth"
	"github.com/4nthem/State/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
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

func NewUserController() (userControl *UserController, err error) {
	if err = auth.InitTokens(); err != nil {
		return
	}
	userControl = &UserController{}
	return
}

func (UserController UserController) Home(rend render.Render, req *http.Request, p martini.Params) {
	rend.JSON(200, map[string]interface{}{"params": p})
}

func (userController UserController) GetUser(rend render.Render, req *http.Request, p martini.Params) {
	rend.JSON(200, map[string]interface{}{"specificUser": p})
}

// gets an individual user
func (userController UserController) GetAllUsers(rend render.Render, req *http.Request, p martini.Params) {

	// get user
	user := models.User{
		// Id:     params.ByName("id"),
		Id:    bson.NewObjectId(),
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	rend.JSON(200, user)

}

// creates a new user
func (userController UserController) CreateUser(rend render.Render, req *http.Request, p martini.Params) {

	// create user
	user := models.User{
		// Id:     params.ByName("id"),
		Id:    bson.NewObjectId(),
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	theToken, _ := auth.NewToken("student", user)

	rend.JSON(200, map[string]interface{}{"NewToken": theToken})

}

// removes an existing user
func (userController UserController) RemoveUser(rend render.Render, req *http.Request, p martini.Params) {

	// delete user

	rend.JSON(200, "deleted User")

}
