package main

import (
	"os"

	"userService/handler"
	"userService/repository"

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

	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
