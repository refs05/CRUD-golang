package controllers

import (
	"crudgoeg/models"
	"time"

	
)

type RequestUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Client_id int    `json:"client_id"`
}

func (req *RequestUser) toModel() models.User {
	return models.User{
		ID:        req.ID,
		Name:      req.Name,
		Age:       req.Age,
		Sex:       req.Sex,
		Client_id: req.Client_id,
	}
}

type ResponseUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Sex       string    `json:"sex"`
	Client_id int       `json:"client_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newResponse(mdlUsers models.User) ResponseUser {
	return ResponseUser{
		ID:        mdlUsers.ID,
		Name:      mdlUsers.Name,
		Age:       mdlUsers.Age,
		Sex:       mdlUsers.Sex,
		Client_id: mdlUsers.Client_id,
		CreatedAt: mdlUsers.CreatedAt,
		UpdatedAt: mdlUsers.UpdatedAt,
	}
}

func newResponseArray(mdlUsers []models.User) []ResponseUser {
	var resp []ResponseUser

	for _, value := range mdlUsers {
		resp = append(resp, newResponse(value))
	}

	return resp
}
