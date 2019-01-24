// 个人主页
package controller

import (
	"fmt"
	"net/http"

	"blog/middleware"
	"blog/util"
	"blog/vm"
)

type profile struct{}

func (this profile) registerRoutes() {
	http.HandleFunc("/user", middleware.Auth(profileHandle))
}

func profileHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pUsername := r.Form.Get("username")
	profileVM := vm.ProfileVM{}
	sUsername, _ := util.GetSessionUser(r)
	v, err := profileVM.GetVM(pUsername, sUsername)
	if err != nil {
		msg := fmt.Sprintf("user ( %s ) does not exist", pUsername)
		w.Write([]byte(msg))
		return
	}

	templates["profile.html"].Execute(w, &v)
}
