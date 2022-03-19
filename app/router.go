package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mFsl16/go-restapi-example/controller"
)

func NewRouter(controller controller.TodoController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", controller.Home)
	router.POST("/todo/create", controller.CreateTodo)
	router.PUT("/todo/update/:id", controller.UpdateTodo)
	router.DELETE("/todo/delete/:id", controller.DeleteTodo)
	router.GET("/todo/:id", controller.FindById)
	router.GET("/todo", controller.FindAll)

	return router

}

func NewServer(router *httprouter.Router) *http.Server {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &server
}
