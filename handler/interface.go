package handler

import "github.com/labstack/echo/v4"

type EndpointInterface interface {
	Hello(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	UpdateUserProfile(ctx echo.Context) error
}
