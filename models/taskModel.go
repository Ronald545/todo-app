package Models

import (
	"github.com/kamva/mgm/v3"
)

// Model Definition
type Task struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	Author           string `json:"author" bson:"author"`
}

func NewTask(name string, description string, author string) *Task {
	return &Task{
		Name:        name,
		Description: description,
		Author:      author,
	}
}
