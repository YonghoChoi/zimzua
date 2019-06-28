package main

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type BusinessDay struct {
	Day       string `json:"day" bson:"day"`
	StartTime string `json:"startTime" bson:"startTime"`
	EndTime   string `json:"endTime" bson:"endTime"`
}

type Storage struct {
	ID           string        `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Phone        string        `json:"phone" bson:"phone"`
	Address      string        `json:"address" bson:"address"`
	BusinessDays []BusinessDay `json:"businessDays" bson:"businessDays"`
	Location     Location      `json:"location" bson:"location"`
	Created      string        `json:"created,omitempty" bson:"created,omitempty"`
	Updated      string        `json:"updated,omitempty" bson:"updated,omitempty"`
}
