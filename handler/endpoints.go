package handler

import (
	"net/http"
	"userService/model"
	"userService/model/constant"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (s *Server) Hello(ctx echo.Context) error {
	resp := "Hello World!"
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) CreateUser(ctx echo.Context) error {
	request := new(model.RequestCreateUser)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResponseError{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	if err := s.Validate.Struct(request); err != nil {
		var errMessageList = make(map[string]interface{})

		for _, err := range err.(validator.ValidationErrors) {
			errMessageList[err.Field()] = []string{TranslationFn(s.Translator, err)}
		}

		return ctx.JSON(http.StatusBadRequest, model.ResponseErrorValidation{
			Status:           constant.ERROR,
			ValidationErrors: errMessageList,
		})
	}

	return ctx.JSON(201, request)
}
