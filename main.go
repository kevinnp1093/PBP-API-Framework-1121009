package main

import (
	"github.com/go-martini/martini"

	"github.com/your-username/your-app/controller"
)

func main() {
	m := martini.Classic()

	m.Get("/users/:id", controller.GetUser)
	m.Post("/users", controller.CreateUser)
	m.Put("/users/:id", controller.UpdateUser)
	m.Delete("/users/:id", controller.DeleteUser)

	m.Run()
}
