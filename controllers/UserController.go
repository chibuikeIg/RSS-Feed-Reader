package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/chibuikeIg/Rss_blog/config"
	router "github.com/julienschmidt/httprouter"
)

var tpl *template.Template

var DB *config.Database

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

type UserController struct{}

func NewUserController(DBConn *config.Database) *UserController {
	DB = DBConn
	return &UserController{}

}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Fatalln(err)
	}

}
