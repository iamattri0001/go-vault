package models

type User struct {
	Base               `bson:",inline"`
	Username           string `json:"username" bson:"username"`
	MasterPasswordHash string `json:"-" bson:"master_password_hash"`
}
