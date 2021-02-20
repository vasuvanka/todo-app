package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/vasuvanka/todo-app/backend/controllers"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/shared"
)

//GetUser - get user handler
func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	paramID := params.ByName("id")
	ID := claims["id"].(string)
	if ID != paramID {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "not a valid id"})
		return
	}
	user, err := controllers.GetUserByID(ID)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, user)
}
