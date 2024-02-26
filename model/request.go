package model

type RequestCreateUser struct {
	FullName string `json:"full_name" validate:"min=3,max=60"`
	Credentials
}

type RequestUpdateUser struct {
	FullName    *string `json:"full_name,omitempty" validate:"omitempty,min=3,max=60"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=13,startswith=+62"`
}

type Credentials struct {
	PhoneNumber string `json:"phone_number" validate:"min=10,max=13,startswith=+62"`
	Password    string `json:"password" validate:"min=6,max=64,hasupper,hasnumber,hasspecial"`
}
