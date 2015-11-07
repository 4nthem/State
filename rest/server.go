package rest

import (
	"github.com/4nthem/State/controllers"
	log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
)

func StartServer() {

	m := martini.Classic()

	user := controllers.NewUserController()

	m.Get("/", user.Home)

	m.Group("/user", func(r martini.Router) {
		r.Get("", user.Home)           // Don't like this. Its hacky but didn't see better solution
		r.Get("/new", user.CreateUser) //Obviously this will be a post
		r.Get("/all", user.GetAllUsers)
		r.Get("/remove", user.RemoveUser)
		r.Get("/(?P<name>[a-zA-Z]+)", user.GetUser)
	}, controllers.TestMiddleware1, controllers.TestMiddleware2)

	log.Info("Starting Server on port 3000")
	m.Run()

}
