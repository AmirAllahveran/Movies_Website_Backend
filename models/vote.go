package models

type Vote struct {
	Id       	int 			`json:"id"`
	User     	User            `json:"user,omitempty"`
	Rating 		complex64       `json:"rating,omitempty" validate:"required"`
	MovieID    	int             `json:"movieID,omitempty" validate:"required"`
}