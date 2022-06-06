package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	router "github.com/julienschmidt/httprouter"
)

type FeedSettingController struct{}

func NewFeedSettingController(DBConn *config.Database) *FeedSettingController {

	DB = DBConn

	return &FeedSettingController{}

}

func (fsc FeedSettingController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {
	middleware.Auth(w, r)

	View(w, "setting.html", nil)
	return
}

func (fsc FeedSettingController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	validation := fsc.ValidateRequest(r)

	if len(validation) != 0 {

		json.NewEncoder(w).Encode(validation)

		return
	}

	json.NewEncoder(w).Encode([]string{})

	return

}

func (fsc FeedSettingController) ValidateRequest(r *http.Request) map[string]string {

	errors := make(map[string]string)

	if r.FormValue("summary_length") == "" {

		errors["summary_length"] = "The feed summary length field is required."

	} else if _, err := strconv.Atoi(r.FormValue("summary_length")); err != nil {

		errors["summary_length"] = "The value entered is invalid."

	}

	if r.FormValue("polling_frequency") == "" {

		errors["polling_frequency"] = "The polling frequency field is required."

	} else if _, err := strconv.Atoi(r.FormValue("polling_frequency")); err != nil {

		errors["polling_frequency"] = "The value entered is invalid."

	}

	return errors
}
