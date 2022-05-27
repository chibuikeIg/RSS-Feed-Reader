package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	router "github.com/julienschmidt/httprouter"
)

type LoginController struct{}

func NewLoginController(DBConn *config.Database) *LoginController {

	DB = DBConn

	return &LoginController{}
}

func (lc LoginController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Guest(w, r)

	View(w, "login.html", nil)
}
