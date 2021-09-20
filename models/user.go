package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Client_id int `json:"client_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated-at"`
	Deleted_at time.Time `json:"deleted_at"`
}

