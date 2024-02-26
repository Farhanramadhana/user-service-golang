package handler

import (
	"context"
	"net/http"
	"time"
	"userService/model"
	"userService/model/constant"
	"userService/repository"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	errMessageList, isError := s.validate(request)
	if isError {
		return ctx.JSON(http.StatusBadRequest, model.ResponseErrorValidation{
			Status:           constant.ERROR,
			ValidationErrors: errMessageList,
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// get user by phone number, if exist return error
	existingUser, _ := s.Repository.GetUserByPhone(context.TODO(), request.PhoneNumber)
	if existingUser.Id != 0 {
		return ctx.JSON(400, model.ResponseError{
			Status:  constant.ERROR,
			Message: "phone number already exist",
		})
	}

	// register user, phone number is unique
	newUser := repository.UserTable{
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
		Password:    string(hash),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = s.Repository.CreateUser(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return ctx.JSON(201, request)
}

func (s *Server) Login(ctx echo.Context) error {
	request := new(model.Credentials)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResponseError{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	errMessageList, isError := s.validate(request)
	if isError {
		return ctx.JSON(http.StatusBadRequest, model.ResponseErrorValidation{
			Status:           constant.ERROR,
			ValidationErrors: errMessageList,
		})
	}

	userData, err := s.Repository.GetUserByPhone(context.TODO(), request.PhoneNumber)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResponseError{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	isValid := s.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return ctx.JSON(http.StatusBadRequest, model.ResponseError{
			Status:  constant.ERROR,
			Message: "wrong password",
		})
	}

	token, err := s.Helper.GenerateToken(userData.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResponseError{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.ResponseToken{
		Token: token,
	})
}

func (s *Server) verifyPassword(password, passwordHash string) bool {
	byteHash := []byte(passwordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	
	return err == nil
}

func (s *Server) validate(request interface{}) (errMessageList map[string]interface{}, isError bool) {
	if err := s.Validate.Struct(request); err != nil {
		errMessageList = make(map[string]interface{})
		isError = true
		for _, err := range err.(validator.ValidationErrors) {
			errMessageList[err.Field()] = []string{TranslationFn(s.Translator, err)}
		}
		return
	}

	return
}
