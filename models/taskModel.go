package taskModel

import (
  "github.com/kamva/mgm/v3"
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
