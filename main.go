package main

import (
	"fmt"

	"github.com/vasuvanka/todo-app/backend/config"
	"github.com/vasuvanka/todo-app/backend/routes"
)

func main() {
	config := config.New()
	config.Init()
	fmt.Println(config)
	server := routes.NewServer(config)
	server.Init()
	server.Bootstrap()
}
