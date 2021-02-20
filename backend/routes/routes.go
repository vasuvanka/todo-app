package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/vasuvanka/todo-app/backend/config"
	"github.com/vasuvanka/todo-app/backend/controllers"
	"github.com/vasuvanka/todo-app/backend/handlers"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/services"
	"github.com/vasuvanka/todo-app/backend/shared"
)

//Server config
type Server struct {
	Config  *config.Config
	Router  *httprouter.Router
	service services.Service
}

//NewServer - new server instance
func NewServer(config *config.Config) *Server {
	return &Server{Config: config}
}

// Init - server initialization
func (s *Server) Init() {

	if err := s.service.Connect(s.Config.DatabaseURI); err != nil {
		log.Fatal(err.Error())
	}
	if err := s.service.Mongo.Index(); err != nil {
		log.Fatal(err.Error())
	}
	s.Router = httprouter.New()
	s.Router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        header := w.Header()
        header.Set("Access-Control-Allow-Methods", "*")
        header.Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNoContent)
	})
	s.Router.GET("/api", handlers.Check)

	// register routes
	controllers.New(&s.service)
	GeneralRoutes(s)
	DirectoryRoutes(s)
	UserRoutes(s)
	TodoRoutes(s)
	

	s.Router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Println(r.RequestURI, i)
		shared.SendError(w, models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	// 404
	s.Router.NotFound = http.NotFoundHandler()
	s.Router.RedirectTrailingSlash = true
}

// Bootstrap - Start the server
func (s *Server) Bootstrap() {
	// s.Router.ServeFiles("/*",http.Dir(s.Config.FEPath))
	fmt.Println("Server listening on " + s.Config.Port)
	handler := cors.AllowAll().Handler(s.Router)
	log.Fatal(http.ListenAndServe(":"+s.Config.Port, handler))
}
