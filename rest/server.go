package rest


import (
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	// "gopkg.in/mgo.v2"
	"github.com/4nthem/State/controllers"
	"github.com/4nthem/State/config"
	"net/http"
)


type RestServer struct {
	port int
}



func NewRestServer() (r *RestServer){
	r = &RestServer{port: 3000}
	return
}


func (r RestServer) StartServer() {
	// Instantiate a new router
	router := httprouter.New()

	// Intialize Configurations
	configuration, err := config.GetConfig()
	if err != nil {
		return
	}

	log.Debug("Confugration: ", configuration)

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


func (r *RestServer) Close() (err error) {
	return
}