package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userService/helpers"
	"userService/model"
	"userService/repository"

	eng "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"full_name":"John Doe","phone_number":"+62812345612","password":"Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	echoContext := echo.New()
	mockCtx := echoContext.NewContext(req, rec)

	en := eng.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")
	validate := NewValidator(translator)
	server := Server{
		Repository: mockRepo,
		Validate:   validate,
	}

	// Mocking the Repository.GetUserByPhone method to return an existing user if exist
	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), gomock.Any()).Return(repository.UserTable{Id: 0}, nil)

	// Mocking the Repository.CreateUser method to return a new user ID
	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(int(1), nil)

	err := server.CreateUser(mockCtx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var userResponse model.ResponseID
		err := json.Unmarshal(rec.Body.Bytes(), &userResponse)
		if assert.NoError(t, err) {
			assert.Equal(t, model.ResponseID{ID: 1}, userResponse)
		}
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	mockHelper := helpers.NewMockHelperInterface(ctrl)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone_number":"+62812345612","password":"Password123@"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	echoContext := echo.New()
	mockCtx := echoContext.NewContext(req, rec)

	en := eng.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")
	validate := NewValidator(translator)
	server := Server{
		Repository: mockRepo,
		Validate:   validate,
		Helper:     mockHelper,
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("Password123@"), bcrypt.MinCost)
	// Mocking the Repository.GetUserByPhone method to return an existing user if exist
	mockRepo.EXPECT().GetUserByPhone(gomock.Any(), gomock.Any()).Return(repository.UserTable{
		Id:          1,
		FullName:    "John Doe",
		PhoneNumber: "+62812345612",
		Password:    string(hash),
	}, nil)

	// Mocking the GenerateToken method to return a predefined token
	mockHelper.EXPECT().GenerateToken(gomock.Any()).Return(gomock.Any().String(), nil)
	mockRepo.EXPECT().UpsertLoginLog(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	err := server.Login(mockCtx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var loginResponse model.ResponseToken
		err := json.Unmarshal(rec.Body.Bytes(), &loginResponse)
		if assert.NoError(t, err) {
			assert.Equal(t, model.ResponseToken{Token: gomock.Any().String()}, loginResponse)
		}
	}
}
