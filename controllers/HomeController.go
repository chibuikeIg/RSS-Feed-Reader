package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	router "github.com/julienschmidt/httprouter"
)

type HomeController struct{}

func NewHomeController(DBConn *config.Database) *HomeController {

	DB = DBConn

	return &HomeController{}

}

func (hc HomeController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	View(w, "index.html", nil)

}
