package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

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

func (fc FollowingController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	validation := fc.ValidateAddFeedRequest(r)

	if len(validation) != 0 {

		json.NewEncoder(w).Encode(validation)

		return
	}

	json.NewEncoder(w).Encode([]int{1, 2, 4, 5})

	return
}

func (fc FollowingController) ValidateAddFeedRequest(r *http.Request) map[string]string {

	errors := make(map[string]string)

	if r.FormValue("rss_link") == "" {

		errors["rss_link"] = "The rss link field is required."

	} else {

		_, err := url.ParseRequestURI(r.FormValue("rss_link"))

		if err != nil {

			errors["rss_link"] = "Invalid url provided."

		}

	}

	return errors
}
