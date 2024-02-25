package model

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseErrorValidation struct {
	Status           string            `json:"status"`
	ValidationErrors []ValidationError `json:"errors"`
}

type ValidationError struct {
	FieldName    string `json:"field_name"`
	ErrorMessage string `json:"error"`
}
