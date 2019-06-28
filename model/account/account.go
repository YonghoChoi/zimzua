package account

type Account struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Phone     string `json:"phone" bson:"phone"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	LoginType string `json:"login_type" bson:"login_type"`
	Token     string `json:"token" bson:"token"`
	Created   string `json:"created,omitempty" bson:"created,omitempty"`
	Updated   string `json:"updated,omitempty" bson:"updated,omitempty"`
}
