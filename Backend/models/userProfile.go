package models

type UserProfile struct {
	Username  string `bson:"username" json:"username"`
	FirstName string `bson:"firstname" json:"firstname"`
	LastName  string `bson:"lastname" json:"lastname"`
	Email     string `bson:"email" json:"email"`
}
