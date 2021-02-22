package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"github.com/vasuvanka/todo-app/backend/controllers"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/shared"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
)

//GetUserTodos - get user todos handler
func GetUserTodos(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	skipLiteral := r.URL.Query().Get("skip")
	limitLiteral := r.URL.Query().Get("limit")
	if skipLiteral == "" {
		skipLiteral = "0"
	}
	if limitLiteral == "" {
		limitLiteral = "10"
	}
	skip, err := strconv.Atoi(skipLiteral)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	limit, err := strconv.Atoi(skipLiteral)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	ID := claims["id"].(string)

	todos, err := controllers.GetUserTodos(ID, params.ByName("id"), skip, limit)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, todos)
}

//GetTodo - get todo handler
func GetTodo(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	ID := params.ByName("id")
	if ID == "" || !bson.IsObjectIdHex(ID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid todo id"})
		return
	}
	todo, err := controllers.GetTodoByID(ID)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, todo)
}

//DeleteTodo - delete todo handler
func DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID := params.ByName("id")
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	if ID == "" || !bson.IsObjectIdHex(ID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid todo id"})
		return
	}
	err := controllers.DeleteTodoByID(ID, claims["id"].(string))
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, models.Response{
		Message: "removed",
	})
}

//UpdateTodo - update todo handler
func UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	ID := params.ByName("id")
	if ID == "" || !bson.IsObjectIdHex(ID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid todo id"})
		return
	}
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	todo.ID = bson.ObjectIdHex(ID)
	todo.CreatedBy = claims["id"].(string)
	dbTodo, err := controllers.UpdateTodo(todo)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, dbTodo)
}

//CreateTodo - create a todo
func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	// body validation
	valid := vd.Validate(todo)
	if valid != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: valid.Error() })
		return
	}

	todo.CreatedBy = claims["id"].(string)
	dbTodo, err := controllers.CreateTodo(todo)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusCreated, dbTodo)
}

//ShareTodo share todo with other user
func ShareTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	fromUserID := claims["id"].(string)
	todoID := params.ByName("id")
	if todoID == "" || !bson.IsObjectIdHex(todoID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid todo id"})
		return
	}
	var share models.ShareTodo
	if err := json.NewDecoder(r.Body).Decode(&share); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	valid := vd.Validate(share)
	if valid != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: valid.Error() })
		return
	}

	err := controllers.ShareTodo(todoID, fromUserID, share.Email)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, models.Response{
		Message: "shared",
	})
}
