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

func (uc UserController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	validation := uc.ValidateLoginRequest(r)

	if validation != nil {

	}
}

func (uc UserController) ValidateLoginRequest(r *http.Request) map[string]string {

	return map[string]string{}

}
