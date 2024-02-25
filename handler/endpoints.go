package handler

import (
	"net/http"
	"userService/model"
	"userService/model/constant"
	"userService/repository"

	"github.com/labstack/echo/v4"
)

func (s *Server) Hello(ctx echo.Context) error {
	resp := "Hello World!"
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) CreateUser(ctx echo.Context) error {
	request := new(repository.RequestCreateUser)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResponseError{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	// add validation
	return ctx.JSON(201, request)
}
