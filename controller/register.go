package controller

import (
	"log"
	"net/http"

	"blog/model"
	"blog/util"
	"blog/validate"
	"blog/vm"
)

type register struct{}

func (this register) registerRoutes() {
	http.HandleFunc("/register", registerHandle)
}

func registerHandle(w http.ResponseWriter, r *http.Request) {
	registerVM := vm.RegisterVM{}
	v := registerVM.GetVM()
	if r.Method == http.MethodGet {
		templates["register.html"].Execute(w, &v)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		if errCheck := validate.CheckRegister(username, email, pwd1, pwd2); len(errCheck) > 0 {
			for _, errMsg := range errCheck {
				v.AddError(errMsg)
			}
		}

		log.Println(v.Errs)
		if len(v.Errs) > 0 {
			templates["register.html"].Execute(w, &v)
			return
		}

		if err := model.AddUser(username, pwd1, email); err != nil {
			log.Println("add User error:", err)
			w.Write([]byte("Error insert database"))
			return
		}

		util.SetSessionUser(w, r, username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
