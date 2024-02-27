package handler

import "github.com/labstack/echo/v4"

type EndpointInterface interface {
	CreateUser(ctx echo.Context) error
	Login(ctx echo.Context) error
	GetUserProfile(ctx echo.Context) error
	UpdateUserProfile(ctx echo.Context) error
}
