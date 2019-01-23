package controller

import (
	"net/http"

	"blog/middleware"
	"blog/util"
	"blog/vm"
)

type index struct{}

func (this index) registerRoutes() {
	http.HandleFunc("/", middleware.Auth(indexHandle))
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	indexVM := vm.IndexVM{}
	username, _ := util.GetSessionUser(r)
	v := indexVM.GetVM(username)
	templates["index.html"].Execute(w, &v)
}
