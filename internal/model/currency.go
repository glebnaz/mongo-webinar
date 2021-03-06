package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Currency struct {
	ID       primitive.ObjectID `json:"ID" bson:"_id"`
	NumCode  string             `json:"NumCode" bson:"num_code"`
	CharCode string             `json:"CharCode" bson:"char_code"`
	Nominal  int                `json:"Nominal" bson:"nominal"`
	Name     string             `json:"Name" bson:"name"`
	Value    float64            `json:"Value" bson:"value"`
	Previous float64            `json:"Previous" bson:"previous"`
	Date     int64              `json:"Date" bson:"date"`
}
