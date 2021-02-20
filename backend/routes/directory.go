package routes

import (
	"github.com/vasuvanka/todo-app/backend/handlers"
	"github.com/vasuvanka/todo-app/backend/middlewares"
)

//DirectoryRoutes - directory routes
func DirectoryRoutes(s *Server) {
	s.Router.POST("/api/directories", middlewares.JwtValidatorWrap(handlers.CreateDir))
	s.Router.GET("/api/directories/:id", middlewares.JwtValidatorWrap(handlers.GetDirByID))
	s.Router.PUT("/api/directories/:id", middlewares.JwtValidatorWrap(handlers.UpdateDir))
	s.Router.DELETE("/api/directories/:id", middlewares.JwtValidatorWrap(handlers.DeleteDir))
	s.Router.GET("/api/directories/:id/todos", middlewares.JwtValidatorWrap(handlers.GetUserTodos))
	s.Router.GET("/api/directories/:id/subdirs", middlewares.JwtValidatorWrap(handlers.GetUserDirs))

}
