package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/4nthem/State/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Get a UserController instance
	// userController := controllers.NewUserController(initDb())
	userController := controllers.NewUserController()

	// Get all user resources
	router.GET("/users", userController.GetUsers)

	// Create a new user
	router.POST("/user", userController.CreateUser)

	// Remove an existing user
	router.DELETE("/user/:id", userController.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", router)
}

// getSession creates a new mongo session and panics if connection error occurs
func initDb() *mgo.Session {
	// Connect to our local mongo
	session, err := mgo.Dial("mongodb://<uri goes here>")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return session
}