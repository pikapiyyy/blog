package controller

import (
	"net/http"

	"blog/middleware"
	"blog/util"
	"blog/validate"
	"blog/vm"
)

type login struct{}

func (this login) registerRoutes() {
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/logout", middleware.Auth(logoutHandle))
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	loginVM := vm.LoginVM{}
	v := loginVM.GetVM()
	if r.Method == http.MethodGet {
		templates["login.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// 用户名密码验证
		if errCheck := validate.CheckLogin(username, password); len(errCheck) > 0 {
			for _, errMsg := range errCheck {
				v.AddError(errMsg)
			}
		}

		if len(v.Errs) > 0 {
			templates["login.html"].Execute(w, &v)
		} else {
			util.SetSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	util.ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
