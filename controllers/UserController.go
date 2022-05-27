package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	router "github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController(DBConn *config.Database) *UserController {
	DB = DBConn
	return &UserController{}
}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	View(w, "index.html", nil)

}
