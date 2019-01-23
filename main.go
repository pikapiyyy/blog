package main

import (
	"net/http"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"blog/controller"
	"blog/model"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	controller.Startup()
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
