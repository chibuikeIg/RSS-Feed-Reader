package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	router "github.com/julienschmidt/httprouter"
)

type FollowingController struct{}

func NewFollowingController(DBConn *config.Database) *FollowingController {

	DB = DBConn

	return &FollowingController{}

}

func (fc FollowingController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	View(w, "index.html", nil)

}

func (fc FollowingController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {
	middleware.Auth(w, r)

	View(w, "create.html", nil)
}
