package model

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
