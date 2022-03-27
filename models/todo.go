package models

import (
	"time"
)

type Todo struct{
	ID uint `json:"id"`
	User_id uint `json:"user_id"`
	Tag_id uint `json:"tag_id"`
	Description string `json:"description"`
	Note string `json:"note"`
	//0->pending   1->completed  2->in-progress
	Status uint `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deadline time.Time `json:"deadline"`
}