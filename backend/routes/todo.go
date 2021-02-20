package routes

import (
	"github.com/vasuvanka/todo-app/backend/handlers"
	"github.com/vasuvanka/todo-app/backend/middlewares"
)

//TodoRoutes - todo routes
func TodoRoutes(s *Server) {
	s.Router.GET("/api/todos/:id", middlewares.JwtValidatorWrap(handlers.GetTodo))
	s.Router.PUT("/api/todos/:id", middlewares.JwtValidatorWrap(handlers.UpdateTodo))
	s.Router.DELETE("/api/todos/:id", middlewares.JwtValidatorWrap(handlers.DeleteTodo))
	s.Router.POST("/api/todos", middlewares.JwtValidatorWrap(handlers.CreateTodo))
	s.Router.POST("/api/todos/:id/share", middlewares.JwtValidatorWrap(handlers.ShareTodo))
}
