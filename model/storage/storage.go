package storage

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Storage struct {
	ID       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Phone    string   `json:"phone" bson:"phone"`
	Address  string   `json:"address" bson:"address"`
	Location Location `json:"location" bson:"location"`
	Created  string   `json:"created,omitempty" bson:"created,omitempty"`
	Updated  string   `json:"updated,omitempty" bson:"updated,omitempty"`
}
