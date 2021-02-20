package controllers

import "github.com/vasuvanka/todo-app/backend/services"

// New - get serivces in controllers
func New(service *services.Service) {
	NewUserController(&service.Db)
	NewTodoController(&service.Db)
	NewDirectoryController(&service.Db)
}
