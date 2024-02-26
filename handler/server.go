package handler

import (
	"userService/helpers"
	"userService/middleware"
	"userService/repository"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

type Server struct {
	Translator ut.Translator
	Validate   *validator.Validate
	Helper     helpers.HelperInterface
	Middleware middleware.MiddlewareInterface
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Translator ut.Translator
	Validate   *validator.Validate
	Helper     helpers.HelperInterface
	Middleware middleware.MiddlewareInterface
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Translator: opts.Translator,
		Validate:   opts.Validate,
		Helper:     opts.Helper,
		Middleware: opts.Middleware,
		Repository: opts.Repository,
	}
}
