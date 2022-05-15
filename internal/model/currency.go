package model

type Currency struct {
	ID       string  `json:"ID" bson:"id"`
	NumCode  string  `json:"NumCode" bson:"num_code"`
	CharCode string  `json:"CharCode" bson:"char_code"`
	Nominal  int     `json:"Nominal" bson:"nominal"`
	Name     string  `json:"Name" bson:"name"`
	Value    float64 `json:"Value" bson:"value"`
	Previous float64 `json:"Previous" bson:"previous"`
	Date     int64   `json:"Date" bson:"date"`
}
