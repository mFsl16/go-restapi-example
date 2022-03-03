package app

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mFsl16/go-restapi-example/utils"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(rw)
		err := encoder.Encode("Welcome To My API")
		utils.PanicIfErr(err)
	})

	return router

}

func NewServer(router *httprouter.Router) *http.Server {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &server
}
