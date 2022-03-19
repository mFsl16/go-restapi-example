package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mFsl16/go-restapi-example/app"
	"github.com/mFsl16/go-restapi-example/utils"
)

func main() {
	server := app.NewApp()

	fmt.Println("Server run on port: 3000")

	err := server.ListenAndServe()
	utils.PanicIfErr(err)
}
