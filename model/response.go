package model

import "time"

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseErrorValidation struct {
	Status           string      `json:"status"`
	ValidationErrors interface{} `json:"errors"`
}

type ResponseToken struct {
	Token string `json:"token"`
}

type ResponseGetUser struct {
	Id          int       `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
