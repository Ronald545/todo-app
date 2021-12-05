package Handlers

import (
	"os"
	"time"

	"github.com/Ronald545/todo-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func FindAllUsers(c *fiber.Ctx) error {
	result := []Models.User{}
	err := mgm.Coll(&Models.User{}).SimpleFind(&result, bson.M{})
	if err != nil {
		return respond(c, 400, "an error occured while trying to find users")
	}
	c.SendStatus(200)
	return c.JSON(result)
}

func LoginUser(c *fiber.Ctx) error {
	u := new(Models.User)
	if err := c.BodyParser(u); err != nil {
		return respond(c, 400, "an error occured while parsing json body")
	}

	user := &Models.User{}
	coll := mgm.Coll(user)
	result := []Models.User{}

	err := coll.SimpleFind(&result, bson.M{"username": u.Username})

	if err != nil {
		return respond(c, 500, "error while searching for user")
	}

	if len(result) == 0 {
		return respond(c, 404, "no users found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(u.Password))

	if err != nil {
		return respond(c, 400, "password incorrect, you are not authorized to delete this account")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":  result[0].ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func CreateUser(c *fiber.Ctx) error {
	u := new(Models.User)
	if err := c.BodyParser(u); err != nil {
		return respond(c, 400, "an error occured while parsing json body")
	}
	user, err := Models.NewUser(u.Username, u.Password)

	if err != nil {
		return respond(c, 500, err.Error())
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		return respond(c, 500, "an error occured while saving the user")
	}

	return respond(c, 200, "user sucessfully registered")
}

func DeleteUser(c *fiber.Ctx) error {
	u := new(Models.User)
	if err := c.BodyParser(u); err != nil {
		return respond(c, 400, "error occured while parsing json")
	}

	user := &Models.User{}
	coll := mgm.Coll(user)
	result := []Models.User{}

	err := coll.SimpleFind(&result, bson.M{"username": u.Username})

	if err != nil {
		return respond(c, 500, "error while searching for user")
	}

	if len(result) == 0 {
		return respond(c, 404, "no users found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(u.Password))

	if err != nil {
		return respond(c, 400, "password incorrect, you are not authorized to delete this account")
	}

	if err := coll.Delete(&result[0]); err != nil {
		return respond(c, 500, "an error occured while deleting this user")
	}

	return respond(c, 200, "user deleted sucessfully")
}
