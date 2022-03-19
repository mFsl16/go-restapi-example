package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	Home(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
