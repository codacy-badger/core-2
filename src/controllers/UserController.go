package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pravinba9495/iborg-core/src/middleware"
	"github.com/pravinba9495/iborg-core/src/models"
	"github.com/pravinba9495/iborg-core/src/services"
)

// UserController represents the entry point for the User API
func UserController(res http.ResponseWriter, req *http.Request) {
	log.Println("User API")

	switch req.Method {
	case "GET":
		getUser(res, req)
		break
	case "POST":
		saveUser(res, req)
		break
	case "PATCH":
		updateUser(res, req)
		break
	default:
		res.WriteHeader(405)
		break
	}
}

func getUser(res http.ResponseWriter, req *http.Request) {
	log.Println("Get User")
	middleware.VerifyToken(res, req)
	res.Write([]byte("OK"))
}

func saveUser(res http.ResponseWriter, req *http.Request) {
	log.Println("Save User")

	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(services.SaveUser(user))

}

func updateUser(res http.ResponseWriter, req *http.Request) {
	log.Println("Update User")
	middleware.VerifyToken(res, req)
	res.Write([]byte("OK"))
}
