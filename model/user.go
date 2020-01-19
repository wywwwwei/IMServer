package model

type User struct {
	UserID string`json:"user" binding:"-"`
	Password string `json:"password" binding:"required"`
	Name string `json:"name" binding:"required"`
	Sex string `json:"sex" binding:"required"`
	Email string `json:"email" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type Friend struct{
	Id string `json:"user"`
	Username string `json:"name"`
}