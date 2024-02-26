package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
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
		return ctx.JSON(http.StatusInternalServerError, model.ResponseMessage{
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
		return ctx.JSON(400, model.ResponseMessage{
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

	id, err := s.Repository.CreateUser(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return ctx.JSON(201, model.ResponseID{
		ID: id,
	})
}

func (s *Server) Login(ctx echo.Context) error {
	request := new(model.Credentials)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResponseMessage{
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
		return ctx.JSON(http.StatusBadRequest, model.ResponseMessage{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	isValid := s.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return ctx.JSON(http.StatusBadRequest, model.ResponseMessage{
			Status:  constant.ERROR,
			Message: "wrong password",
		})
	}

	token, err := s.Helper.GenerateToken(userData.Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResponseMessage{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	now := time.Now().UTC()
	err = s.Repository.UpsertLoginLog(context.TODO(), userData.Id, now)
	if err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(http.StatusOK, model.ResponseToken{
		Token: token,
	})
}

func (s *Server) GetUserProfile(ctx echo.Context) error {
	idStr := ctx.Get("user_id").(string)
	id, _ := strconv.Atoi(idStr)

	userData, err := s.Repository.GetUserByID(context.TODO(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ResponseMessage{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	response := model.ResponseGetUser{
		Id:          userData.Id,
		FullName:    userData.FullName,
		PhoneNumber: userData.PhoneNumber,
		CreatedAt:   userData.CreatedAt,
		UpdatedAt:   userData.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response)
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
