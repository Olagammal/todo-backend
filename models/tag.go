package models

import (
	"time"
)

type Tag struct {
	Id       uint   `json:"id"`
	Tagname     string `json:"tag_name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}