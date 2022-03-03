//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/google/wire"
)

func NewApp() *http.Server {
	wire.Build(
		NewRouter,
		NewServer,
	)

	return nil
}
