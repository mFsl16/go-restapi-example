//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/mFsl16/go-restapi-example/controller"
	"github.com/mFsl16/go-restapi-example/repository"
	"github.com/mFsl16/go-restapi-example/service"
)

func NewApp() *http.Server {
	wire.Build(
		NewRouter,
		NewServer,
		repository.NewDB,
		repository.NewTodoRepository,
		service.NewTodoService,
		controller.NewTodoController,
		resty.New,
		controller.NewPostController,
		service.NewHttpClientService,
	)

	return nil
}
