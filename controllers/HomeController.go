package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	router "github.com/julienschmidt/httprouter"
)

type HomeController struct{}

func NewHomeController(DBConn *config.Database) *HomeController {

	DB = DBConn

	return &HomeController{}

}

func (hc HomeController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	http.Redirect(w, r, "/following", http.StatusSeeOther)

	return

}
