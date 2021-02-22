package handlers

import (
	"encoding/json"
	"net/http"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"github.com/vasuvanka/todo-app/backend/controllers"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/shared"
)

//GetUserDirs - get user dirs handler
func GetUserDirs(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	parentID := params.ByName("id")
	if !bson.IsObjectIdHex(parentID) {
		if parentID != "0" {
			shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid parent id"})
			return
		}
	}
	dirs, err := controllers.GetUserDirs(claims["id"].(string), parentID)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, dirs)
}

//GetDirByID - get dir by id handler
func GetDirByID(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	dirID := params.ByName("id")
	if !bson.IsObjectIdHex(dirID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "Invalid id"})
		return
	}
	dirs, err := controllers.GetDirByID(dirID)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, dirs)
}

//DeleteDir - delete dir handler
func DeleteDir(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	ID := params.ByName("id")
	if ID == "" || !bson.IsObjectIdHex(ID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid id"})
		return
	}
	err := controllers.DeleteDir(ID)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, models.Response{
		Message: "removed",
	})
}

//UpdateDir - update dir handler
func UpdateDir(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	ID := params.ByName("id")
	if ID == "" || !bson.IsObjectIdHex(ID) {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid id"})
		return
	}
	var dir models.Directory
	if err := json.NewDecoder(r.Body).Decode(&dir); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	dir.CreatedBy = claims["id"].(string)
	err := controllers.UpdateDir(dir)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusOK, dir)
}

//CreateDir - create a dir
func CreateDir(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	claims := r.Context().Value(shared.KeyClaims).(jwt.MapClaims)
	var dir models.Directory
	if err := json.NewDecoder(r.Body).Decode(&dir); err != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	dir.CreatedBy = claims["id"].(string)

	valid := vd.Validate(dir)
	if valid != nil {
		shared.SendError(w, models.Response{Status: http.StatusBadRequest, Message: valid.Error() })
		return
	}

	dbDir, err := controllers.CreateDir(dir)
	if err != nil {
		shared.SendError(w, models.Response{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	shared.SendJSON(w, http.StatusCreated, dbDir)
}
