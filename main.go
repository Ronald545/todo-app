package main

import (
	"fmt"
	"log"
	"os"

	Handlers "github.com/Ronald545/todo-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ansrivas/fiberprometheus/v2"
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

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	prometheus := fiberprometheus.New("todo-app")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	router(app)

	app.Listen(":5000")
}

func router(app *fiber.App) {

	JWT_SECRET := os.Getenv("JWT_SECRET")

	api := app.Group("/task")

	api.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies("todo-auth")
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %v", cookie))
		return c.Next()
	})

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(JWT_SECRET),
	}))

	auth := app.Group("/auth")

	// tasks
	api.Get("/verify", Handlers.LoginStatus)

	api.Get("/", Handlers.FindTask)

	api.Post("/", Handlers.CreateTask)

	api.Delete("/:id", Handlers.DeleteTask)

	api.Put("/", Handlers.EditTask)

	// auth

	auth.Post("/login", Handlers.LoginUser)

	auth.Post("/signup", Handlers.CreateUser)

	auth.Post("/logout", Handlers.LogOut)

	// dev
	app.Get("/allTasks", Handlers.FindAllTasks)

	app.Get("/allUsers", Handlers.FindAllUsers)

	app.Delete("/users", Handlers.DeleteUser)

	app.Delete("/all", Handlers.DeleteAllUsers)
}
