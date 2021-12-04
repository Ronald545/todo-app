package main

import (
  "github.com/Ronald545/todo-app/handlers"
	"github.com/gofiber/fiber/v2"
  "github.com/kamva/mgm/v3"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  // setup orm and server
  mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017"))

  app := fiber.New()

  router(app)

  app.Listen(":3000")
}

func router(app *fiber.App) {
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Welcome to the TODO API")
  })

  app.Get("/task", taskHandler.FindTask)

  app.Post("/task", taskHandler.CreateTask)

  app.Delete("/task/:id", taskHandler.DeleteTask)

  app.Put("/task", taskHandler.EditTask)
}
