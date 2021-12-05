package Handlers

import (
	"github.com/Ronald545/todo-app/models"
	"github.com/gofiber/fiber/v2"
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

	err = bcrypt.CompareHashAndPassword([]byte(result[0].Password), []byte(u.Password))

	if err != nil {
		return respond(c, 400, "password incorrect, you are not authorized to delete this account")
	}

	if err := coll.Delete(&result[0]); err != nil {
		return respond(c, 500, "an error occured while deleting this user")
	}

	return respond(c, 200, "user deleted sucessfully")
}
