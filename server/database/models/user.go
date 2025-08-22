package models

type User struct {
	Base           `bson:",inline"`
	Username       string `json:"username" bson:"username"`
	PasswordHash   string `json:"-" bson:"password_hash"`
	AuthSalt       string `json:"auth_salt" bson:"auth_salt"`
	EncryptionSalt string `json:"encryption_salt" bson:"encryption_salt"`
}
