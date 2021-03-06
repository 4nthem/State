package controllers

import (
	"net/http"

	// "github.com/4nthem/State/auth"
	"github.com/4nthem/State/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"io/ioutil"
)

type (
	UserController struct {
		collection *mgo.Collection
	}
)

// provides a reference to a UserController with mongo session
func NewUserController(db_collection *mgo.Collection) (userControl *UserController, err error) {  
    // if err = auth.InitTokens(); err != nil {
	// 	return
	// }
    userControl = &UserController{db_collection}
	return
}


// func (UserController UserController) Home(rend render.Render, req *http.Request, p martini.Params) {
// 	rend.JSON(200, map[string]interface{}{"params": p})
// }

func (userController UserController) GetUser(rend render.Render, req *http.Request, p martini.Params) {

	id := p["id"]
	// Grab id
    oid := bson.ObjectIdHex(id)
    user := models.User{}

    // Fetch user
    if err := userController.collection.FindId(oid).One(&user); err != nil {
       
        rend.JSON(404, map[string]interface{}{"message": "Failed to found user", "data": err })
    }

	rend.JSON(200, map[string]interface{}{"message": "Successfully found user", "data": user})
}

// gets an individual user
func (userController UserController) GetAllUsers(rend render.Render, req *http.Request, p martini.Params) {

	var users []models.User

	// get users
	if err := userController.collection.Find(nil).All(&users); err != nil {
       
        rend.JSON(404, map[string]interface{}{"message": "Failed to found users", "data": err })
    }

	rend.JSON(200, map[string]interface{}{"message": "Successfully found users", "data": users})


}

// creates a new user
func (userController UserController) CreateUser(rend render.Render, req *http.Request, p martini.Params) {



	// bind model attributes to user var
    user := models.User{}

    // Populate the user data from the request, should improve using params
    json.NewDecoder(req.Body).Decode(&user)
    user.Id = bson.NewObjectId()


	// theToken, _ := auth.NewToken("student", user)

	if err := userController.collection.Insert(user); err != nil {
       
        rend.JSON(404, map[string]interface{}{"message": "Failed to create user", "data": err })
    }


	// rend.JSON(200, map[string]interface{}{"NewToken": theToken})
	rend.JSON(200, map[string]interface{}{"message": "Successfully created user", "data": user})

}

// removes an existing user
func (userController UserController) DeleteUser(rend render.Render, req *http.Request, p martini.Params) {

	id := p["id"]
	// Grab id
    oid := bson.ObjectIdHex(id)

	// delete user
	if err := userController.collection.RemoveId(oid); err != nil {
       
        rend.JSON(404, map[string]interface{}{"message": "Failed to delete user", "data": err })
    }

	rend.JSON(200, map[string]interface{}{"message": "Successfully deleted user", "data": nil})

}

// update an existing user
func (userController UserController) UpdateUser(rend render.Render, req *http.Request, p martini.Params) {

	// bind model attributes to user var
    // user := models.User{}

	id := p["id"]
	// Grab id
    oid := bson.ObjectIdHex(id)

    // Populate the user data from the request, should improve using params
    data := map[string]interface{}{}
    body, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(body, &data)



	// update user
	if err := userController.collection.UpdateId(oid, bson.M{"$set": data }); err != nil {
       
        rend.JSON(404, map[string]interface{}{"message": "Failed to update user", "data": err })
        return
    }

	rend.JSON(200, map[string]interface{}{"message": "Successfully updated user", "data": nil})

}
