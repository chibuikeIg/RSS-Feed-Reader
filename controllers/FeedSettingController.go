package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	"github.com/chibuikeIg/Rss_blog/models"
	router "github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

type FeedSettingController struct{}

func NewFeedSettingController(DBConn *config.Database) *FeedSettingController {

	DB = DBConn

	return &FeedSettingController{}

}

func (fsc FeedSettingController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {
	middleware.Auth(w, r)

	var settings []models.Setting

	DB.Find("settings").First(&settings)

	View(w, "setting.html", settings)
	return
}

func (fsc FeedSettingController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	validation := fsc.ValidateRequest(r)

	if len(validation) != 0 {

		json.NewEncoder(w).Encode(validation)

		return
	}

	var settings []models.Setting

	DB.Find("settings").First(&settings)

	if len(settings) == 0 {

		if _, err := DB.Collection("settings").InsertOne(context.TODO(), models.Setting{
			Summary_length:    r.FormValue("summary_length"),
			Polling_frequency: r.FormValue("polling_frequency"),
			Last_poll:         time.Now(),
		}); err != nil {
			handleError(err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Technical error occured, please try again."})

			return
		}

	} else {

		filter := bson.D{{"_id", settings[0].Id}}
		update := bson.D{{"$set", models.Setting{
			Summary_length:    r.FormValue("summary_length"),
			Polling_frequency: r.FormValue("polling_frequency"),
			Last_poll:         settings[0].Last_poll,
		}}}

		if _, err := DB.Collection("settings").UpdateOne(context.TODO(), filter, update); err != nil {
			handleError(err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Technical error occured, please try again."})

			return
		}

	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

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
