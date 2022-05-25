package controllers

import (
	"text/template"

	"github.com/chibuikeIg/Rss_blog/config"
)

type LoginController struct {
	View *template.Template
}

func NewLoginController(DBConn *config.Database, view *template.Template) *LoginController {

	DB = DBConn

	return &LoginController{view}
}
