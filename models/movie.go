package models

type Movie struct {
	Id       	int 			`json:"id"`
	Name     	string          `json:"name,omitempty" validate:"required"`
	Description string          `json:"description,omitempty" validate:"required"`
	Rating    	float64       	`json:"rating,omitempty" validate:"max=1,min=0"`
	Comments    []Comment		`json:"comments,omitempty"`
}