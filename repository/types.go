// This file contains types that are used in the repository layer.
package repository

import "time"

type UserTable struct {
	Id          int
	FullName    string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RequestCreateUser struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
