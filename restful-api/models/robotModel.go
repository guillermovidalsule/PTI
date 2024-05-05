package models

import (
	"time"

	"restful-api/structures"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Robot struct {
	ID         primitive.ObjectID `bson:"_id"`
	Robotname  *string            `json:"robotname" validate:"required"`
	Robotidle  bool               `json:"robotidle"`
	Robotstate bool               `json:"robotstate"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Robot_id   string             `json:"robot_id"`
	User_id    string             `json:"user_id"`
	Robot_info structures.Macros  `json:"robot_info"`
	Ruta       []structures.Coord `json:"ruta"`
}
