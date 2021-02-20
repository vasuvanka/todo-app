package routes

import "github.com/vasuvanka/todo-app/backend/handlers"

//GeneralRoutes - general routes
func GeneralRoutes(s *Server) {
	s.Router.POST("/api/signup", handlers.Signup)
	s.Router.POST("/api/login", handlers.Login)
}
