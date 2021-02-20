package main

import (
	"github.com/vasuvanka/todo-app/backend/config"
	"github.com/vasuvanka/todo-app/backend/routes"
)

func main() {
	config := config.New()
	config.FEPath = "./dist/todo"
	config.Init()
	server := routes.NewServer(config)
	server.Init()
	server.Bootstrap()
}