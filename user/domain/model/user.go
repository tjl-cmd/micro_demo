package model

type User struct {
	ID           int64  `json:"id" gorm:"primary_key;not_nul;auto_increment"`
	UserName     string `json:"user_name" gorm:"unique_index;not_null"`
	FirstName    string `json:"first_name"`
	HashPassword string `json:"hash_password"`
}
