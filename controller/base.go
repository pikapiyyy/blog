package controller

import (
	"html/template"

	"blog/util"
)

var (
	indexController    index
	loginController    login
	registerController register
	templates          map[string]*template.Template
)

func init() {
	templates = util.PopulateTemplates()
}

func Startup() {
	indexController.registerRoutes()
	loginController.registerRoutes()
	registerController.registerRoutes()
}
