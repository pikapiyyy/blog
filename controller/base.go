package controller

import (
	"html/template"

	"blog/util"
)

var (
	indexController    index
	loginController    login
	registerController register
	profileController  profile
	templates          map[string]*template.Template
)

func init() {
	templates = util.PopulateTemplates()
}

func Startup() {
	indexController.registerRoutes()
	loginController.registerRoutes()
	registerController.registerRoutes()
	profileController.registerRoutes()
}
