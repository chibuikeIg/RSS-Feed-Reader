package controllers

import (
	"log"
	"net/http"
	"text/template"
	router "github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

type UserController struct {}


func NewUserController() *UserController {
	return &UserController{};
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Fatalln(err)
	}

}