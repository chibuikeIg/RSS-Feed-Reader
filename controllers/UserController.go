package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/chibuikeIg/Rss_blog/config"
	router "github.com/julienschmidt/httprouter"
)

var DB *config.Database

type UserController struct {
	View *template.Template
}

func NewUserController(DBConn *config.Database, view *template.Template) *UserController {
	DB = DBConn

	return &UserController{view}
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	err := uc.View.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Fatalln(err)
	}

}
