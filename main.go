package main

import (
	"log"
	"os"

	"github.com/Ronald545/todo-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  // loading env secrets
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  
  URI := os.Getenv("MONGO_URI")
  DB := os.Getenv("MONGO_DB")
  
  // setup orm and server
  mgm.SetDefaultConfig(nil, DB, options.Client().ApplyURI(URI))

  app := fiber.New()

  router(app)

  app.Listen(":3000")
}

func router(app *fiber.App) {
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Welcome to the TODO API")
  })

  app.Get("/task", Handlers.FindTask)

  app.Post("/task", Handlers.CreateTask)

  app.Delete("/task/:id", Handlers.DeleteTask)

  app.Put("/task", Handlers.EditTask)
}
