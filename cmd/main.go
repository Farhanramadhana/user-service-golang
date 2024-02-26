package main

import (
	"os"

	"userService/handler"
	"userService/helpers"
	"userService/middleware"
	"userService/repository"

	eng "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	s := newServer()

	e.GET("/", func(c echo.Context) error {
		return s.Hello(c)
	})

	e.POST("/user", func(c echo.Context) error {
		return s.CreateUser(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return s.Login(c)
	})

	e.POST("/user/profile", func(c echo.Context) error {
		return s.GetUserProfile(c)
	}, s.Middleware.AuthenticateMiddleware)

	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	en := eng.New()
	uni := ut.New(en, en)
	translator, _ := uni.GetTranslator("en")

	helper := helpers.NewHelper()
	middleware := middleware.NewMiddleware(helper)

	validate := handler.NewValidator(translator)
	opts := handler.NewServerOptions{
		Translator: translator,
		Validate:   validate,
		Helper:     helper,
		Middleware: middleware,
		Repository: repo,
	}

	return handler.NewServer(opts)
}
