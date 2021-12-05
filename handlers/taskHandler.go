package Handlers

import (
	"github.com/Ronald545/todo-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type request struct {
	Id             string `json:"id"`
	NewName        string `json:"newName"`
	NewDescription string `json:"newDescription"`
}

func FindTask(c *fiber.Ctx) error {
	result := []Models.Task{}
	err := mgm.Coll(&Models.Task{}).SimpleFind(&result, bson.M{})
	if err != nil {
		return respond(c, 500, "an error occured while trying to find tasks")
	}

	c.SendStatus(200)
	return c.JSON(result)
}

func CreateTask(c *fiber.Ctx) error {
	// retrieving json
	t := new(Models.Task)
	if err := c.BodyParser(t); err != nil {
		return respond(c, 400, "an error occured while parsing json body")
	}
	// saving into db
	task := Models.NewTask(t.Name, t.Description)
	if err := mgm.Coll(task).Create(task); err != nil {
		return respond(c, 500, "an error occured when saving the task")
	}

	// sucessful response
	return respond(c, 200, "task saved successfully")
}

func DeleteTask(c *fiber.Ctx) error {
	// retrieve id
	task := &Models.Task{}
	coll := mgm.Coll(task)
	if err := coll.FindByID(c.Params("id"), task); err != nil {
		return respond(c, 400, err.Error())
	}

	if err := coll.Delete(task); err != nil {
		return respond(c, 500, "an error occured while deleting this task")
	}

	return respond(c, 200, "task deleted succssfully")
}

func EditTask(c *fiber.Ctx) error {
	// retrieve json body
	task := &Models.Task{}
	coll := mgm.Coll(task)
	b := new(request)

	if err := c.BodyParser(b); err != nil {
		return respond(c, 400, "error while unmarshalling json")
	}

	if err := coll.FindByID(b.Id, task); err != nil {
		return respond(c, 400, "unable to find task")
	}

	task.Name = b.NewName
	task.Description = b.NewDescription

	if err := coll.Update(task); err != nil {
		return respond(c, 500, "error saving task")
	}

	return respond(c, 200, "task sucessfully updated")
}

func respond(ctx *fiber.Ctx, code int, message string) error {
	ctx.SendStatus(code)
	return ctx.SendString(message)
}
