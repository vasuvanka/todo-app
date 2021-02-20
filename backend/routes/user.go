package routes

import (
	"github.com/vasuvanka/todo-app/backend/handlers"
	"github.com/vasuvanka/todo-app/backend/middlewares"
)

//UserRoutes - user routes
func UserRoutes(s *Server) {
	s.Router.GET("/api/users/:id", middlewares.JwtValidatorWrap(handlers.GetUser))
}
