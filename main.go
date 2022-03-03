package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mFsl16/go-restapi-example/app"
	"github.com/mFsl16/go-restapi-example/utils"
)

func main() {
	server := app.NewApp()

	err := server.ListenAndServe()
	utils.PanicIfErr(err)
}
