package rest

import (
	"github.com/4nthem/State/controllers"
	log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func StartServer() {

	m := martini.Classic()

	m.Use(render.Renderer())

	db_session := getSession()
	if db_session == nil {
		return
	}


	user, err := controllers.NewUserController(db_session.C("users"))
	if err != nil {
		log.Error("Error reading mongo db users collection: ", err)
		return
	}

	// m.GET("/", user.Home)  
	//  dashboard or landing page for app   root_path  "/"

	m.Group("/users", func(r martini.Router) {
		r.Get("", user.GetAllUsers)		
		r.Post("", user.CreateUser) 
		// r.Get("/(?P<name>[a-zA-Z]+)", user.GetUser)    //get user's profile
		r.Get("/:id", user.GetUser)
		r.Delete("/(?P<name>[a-zA-Z]+)", user.DeleteUser) //delete user's profile and associations 
		r.Patch("/(?P<name>[a-zA-Z]+)", user.UpdateUser) //update user's profile
	})

	m.Run()

}


func getSession() *mgo.Database {  
    // Connect to our local mongo, gets the state database
    db_session, err := mgo.Dial("mongodb://<user>:<password>@ds051534.mongolab.com:51534/state")

    // Check if connection error, is mongo running?
    if err != nil {
    	log.Error("Error reading mongo db connection: ", err)
		return nil
    }
    
    return db_session.DB("")
}