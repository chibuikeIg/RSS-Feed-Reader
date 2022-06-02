package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	"github.com/chibuikeIg/Rss_blog/models"
	router "github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FollowingController struct{}

func NewFollowingController(DBConn *config.Database) *FollowingController {

	DB = DBConn

	return &FollowingController{}

}

func (fc FollowingController) Index(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	if r.Header.Get("X-Requested-With") == "xmlhttprequest" {

		cursor, err := DB.Collection("posts").Find(DB.Ctx, bson.D{{"deleted_at", nil}}, options.Find().SetSort(bson.D{{"created_at", -1}}))

		if err != nil {
			log.Fatal(err)
		}

		var posts []models.Post

		if err = cursor.All(DB.Ctx, &posts); err != nil {
			log.Fatal(err)
		}

		View(w, "feed-posts.html", posts)

		return
	}

	View(w, "index.html", nil)

	return

}

func (fc FollowingController) Update(w http.ResponseWriter, r *http.Request, ps router.Params) {
	middleware.Auth(w, r)

	status := r.URL.Query().Get("status")

	id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))

	var err error

	if status == "read" {
		_, err = DB.Collection("posts").UpdateOne(DB.Ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"read_at", time.Now()}}}})
	} else if status == "unread" {
		var t time.Time
		_, err = DB.Collection("posts").UpdateOne(DB.Ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"read_at", t}}}})
	}

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to mark post as read."})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

	return
}

func (fc FollowingController) Delete(w http.ResponseWriter, r *http.Request, ps router.Params) {
	middleware.Auth(w, r)

	id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))

	_, err := DB.Collection("posts").UpdateOne(DB.Ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"deleted_at", time.Now()}}}})

	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to delete post."})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

	return
}
