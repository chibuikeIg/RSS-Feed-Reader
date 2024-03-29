package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

		filter := bson.D{{"deleted_at", nil}}

		if r.URL.Query().Get("s_qry") != "" && r.URL.Query().Get("s_qry") != "null" {

			filter = bson.D{{"deleted_at", nil}, {"$text", bson.D{{"$search", r.URL.Query().Get("s_qry")}}}}

		}

		cursor, err := DB.Collection("posts").Find(DB.Ctx, filter, options.Find().SetSort(bson.D{{"created_at", -1}}))

		handleError(err)

		var posts []models.Post

		err = cursor.All(DB.Ctx, &posts)
		handleError(err)

		// get summary settings
		var settings []models.Setting

		DB.Find("settings").First(&settings)

		summary_length := 30

		if len(settings) > 0 {
			summary_length, _ = strconv.Atoi(settings[0].Summary_length)
		}

		View(w, "feed-posts.html", map[string]any{
			"posts":          posts,
			"summary_length": summary_length,
		})

		return
	}

	View(w, "index.html", nil)

	return

}

func (fc FollowingController) LatestPosts(w http.ResponseWriter, r *http.Request, _ router.Params) {
	middleware.Auth(w, r)

	if r.Header.Get("X-Requested-With") == "xmlhttprequest" {

		filter := bson.D{{"deleted_at", nil}}

		cursor, err := DB.Collection("posts").Find(DB.Ctx, filter, options.Find().SetSort(bson.D{{"created_at", -1}}))

		handleError(err)

		var posts []models.Post

		err = cursor.All(DB.Ctx, &posts)

		handleError(err)

		currentPostsCount, _ := strconv.Atoi(r.URL.Query().Get("current_posts_count"))

		if len(posts) > currentPostsCount {

			// get summary settings
			var settings []models.Setting

			DB.Find("settings").First(&settings)

			summary_length := 30

			if len(settings) > 0 {
				summary_length, _ = strconv.Atoi(settings[0].Summary_length)
			}

			View(w, "feed-posts.html", map[string]any{
				"posts":          posts,
				"summary_length": summary_length,
			})

			return

		}

	}

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

	handleError(err)

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

	handleError(err)
	if err == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to delete post."})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

	return
}
