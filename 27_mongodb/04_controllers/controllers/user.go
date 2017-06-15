package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LizzieM2003/golang-web-programming/27_mongodb/04_controllers/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// test via command: curl http://localhost:8080/user/1
	u := models.User{
		Name:   "James",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// send Content-Type, status and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// test via command:
	// curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
	u := models.User{}

	// decode the json in the request body
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "007"

	// marshal the new user into JSON and send it in the response
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// test via command:
	// curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
	// TODO: delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Write code to delete user\n")
}
