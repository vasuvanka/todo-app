package controllers

import (
	"errors"

	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/services/database"
	"github.com/vasuvanka/todo-app/backend/shared"
)

var iUser database.IUser

//NewUserController - will get user interface
func NewUserController(dbUser database.IUser) {
	iUser = dbUser
}

//Login - validates user authentication
func Login(email, password string) (models.User, error) {
	exist, err := iUser.UserExist(email)
	if err != nil {
		return models.User{}, err
	}
	if !exist {
		return models.User{}, errors.New("invalid email or password")
	}
	dbUser, err := iUser.GetUserByUsername(email)
	if err != nil {
		return models.User{}, err
	}
	if !shared.CheckPasswordHash(password, dbUser.Password) {
		return models.User{}, errors.New("invalid email or password")
	}
	return dbUser, err
}

//Singup will register the user into database
func Singup(user models.Signup) error {
	exist, err := iUser.UserExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user email already registered")
	}
	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.Removed = false
	dbUser.Role = shared.RoleUser
	dbUser.Password, err = shared.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = iUser.CreateUser(dbUser)
	return err
}

//GetUserByID - get user by id
func GetUserByID(id string) (models.User, error) {
	return iUser.GetUserByID(id)
}
