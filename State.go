package main

import (
	// Standard library packages
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// Third party packages
	"github.com/4nthem/State/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/4nthem/State/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/4nthem/State/controllers"
)

type config struct {
	Port       string
	Connection string
}

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Intialize Configurations
	configuration, err := getConfig()
	fmt.Println(configuration)

	if err != nil {
		fmt.Println("Error loading configuration: ", err)
		return
	}

	//Get a UserController instance
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

//Function that gets the configuration file and parses into our config stuct
//Notice that the return variable is defined in the signature
func getConfig() (currConfig config, err error) {

	//Open the config file. Defer to closing it when the function goes out of scope
	configFile, err := os.Open("env/config.json")
	defer configFile.Close()

	//If we had an error log it and return our error
	if err != nil {
		fmt.Println("Problem opening config file: ", err)
		return
	}

	// Create a json parser for this file
	jsonParser := json.NewDecoder(configFile)

	//This is a good golang pattern to use:
	// grab the error from the Decode call and if the err is not null log
	if err = jsonParser.Decode(&currConfig); err != nil {
		fmt.Println("Problem parsing the config file: ", err)
		return
	}

	return currConfig, err

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
