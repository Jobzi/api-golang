package model_task

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//taskSchema esquema de la estructura de tareas
type TaskSchema struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty" json:"ID"` 
	Content   string    `json:"content"`
	Important   bool    `json:"important"`
	Date time.Time `bson:"date" json:"date"`
	CreatedAt time.Time `bson:"CreatedAt" json:"CreatedAt"`
	UpdatedAt time.Time `bson:"UpdatedAt" json:"UpdatedAt,omitempty"`
}

//allTasks lista de task
type AllTasks []*TaskSchema
