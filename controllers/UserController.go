package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/chibuikeIg/Rss_blog/config"
	router "github.com/julienschmidt/httprouter"
)

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

func (uc UserController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	validation := uc.ValidateLoginRequest(r)

	if validation != nil {

	}
}

func (uc UserController) ValidateLoginRequest(r *http.Request) map[string]string {

	return map[string]string{}

}
