package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	FindPostById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
