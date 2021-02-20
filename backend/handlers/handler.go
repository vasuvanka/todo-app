package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vasuvanka/todo-app/backend/controllers"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/shared"
)

// Check -  will send ok with 200 status
func Check(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ok")
}

// Signup - create new user
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var singup models.Signup
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&singup); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	// validate user input

	if err := controllers.Singup(singup); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	shared.SendJSON(w, http.StatusCreated, models.Response{
		Message: "user created",
	})
}

// Login - authorize a user
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var login models.Login
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	dbUser, err := controllers.Login(login.Email, login.Password)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	token, err := shared.CreateJWT(dbUser.ID.Hex(), dbUser.Name, dbUser.Role)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, models.LoginResponse{Token: token})
}
