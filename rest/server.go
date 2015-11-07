package rest

import (
	"github.com/4nthem/State/controllers"
	// log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func StartServer() {

	m := martini.Classic()

	m.Use(render.Renderer())

	user, err := controllers.NewUserController()
	if err != nil {
		return
	}

	m.Get("/", user.Home)

	m.Group("/user", func(r martini.Router) {
		r.Get("", user.Home)           // Don't like this. Its hacky but didn't see better solution
		r.Get("/new", user.CreateUser) //Obviously this will be a post
		r.Get("/all", user.GetAllUsers)
		r.Get("/remove", user.RemoveUser)
		r.Get("/(?P<name>[a-zA-Z]+)", user.GetUser)
	})

	m.Run()

}
