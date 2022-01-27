package models

import "time"

type Comment struct {
	Id       	int 			`json:"id"`
	User     	User            `json:"user,omitempty"`
	Comment 	string          `json:"comment_body,omitempty" validate:"required"`
	CreatedAt   time.Time       `json:"createdAt,omitempty"`
	Approved	bool			`json:"approved,omitempty"`
	MovieID		int				`json:"movie_id,omitempty" validate:"required"`
}