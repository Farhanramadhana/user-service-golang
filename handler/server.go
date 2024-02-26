package handler

import (
	"userService/repository"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

type Server struct {
	Translator ut.Translator
	Validate   *validator.Validate
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Translator ut.Translator
	Validate   *validator.Validate
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Translator: opts.Translator,
		Validate:   opts.Validate,
		Repository: opts.Repository,
	}
}
