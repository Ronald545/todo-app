package main

import (
  "fmt"
	"github.com/gofiber/fiber/v2"
  "github.com/kamva/mgm/v3"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
)

// Model Definition 
type Task struct {
  mgm.DefaultModel `bson:",inline"`
  Name string `json:"name" bson:"name"`
  Description string `json:"description" bson:"description"`
}

func NewTask(name string, description string) *Task {
  return &Task{
    Name: name,
    Description: description,
  }
}

func main() {
  // setup orm and server
  mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017"))

  app := fiber.New()

  // routes
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello World")
  })

  app.Get("/task", func(c *fiber.Ctx) error {
    result := []Task{}
    err := mgm.Coll(&Task{}).SimpleFind(&result, bson.M{})
    if err != nil {
      c.SendStatus(400)
      return c.SendString("an error occured while trying to find tasks")
    }

    c.SendStatus(200)
    return c.JSON(result)
  })

  app.Post("/task", func(c *fiber.Ctx) error {
    // retrieving json
    t := new(Task)
    if err := c.BodyParser(t); err != nil {
      fmt.Println(err)
      c.SendStatus(400)
      return c.SendString("an error occured while parsing json body")
    }
    // saving into db
    task := NewTask(t.Name,t.Description)
    if err := mgm.Coll(task).Create(task); err != nil {
      fmt.Println(err)
      c.SendStatus(400)
      return c.SendString("an error occured when saving the task")
    }

    // sucessful response
    c.SendStatus(200)
    return c.SendString("task saved successfully")
  })

  app.Listen(":3000")
}
