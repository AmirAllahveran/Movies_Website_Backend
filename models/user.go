package models

type User struct {
	Id       	int 			`json:"id"`
	Role     	int          	`json:"role,omitempty" validate:"required"`
	Username 	string          `json:"username,omitempty" validate:"required"`
	Password    string          `json:"password,omitempty" validate:"required"`
}