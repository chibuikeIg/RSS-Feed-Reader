package controllers

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	"github.com/chibuikeIg/Rss_blog/models"
	router "github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeedController struct{}

func NewFeedController(DBConn *config.Database) *FeedController {

	DB = DBConn

	return &FeedController{}

}

func (fc FeedController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {
	middleware.Auth(w, r)

	if r.Header.Get("X-Requested-With") == "xmlhttprequest" {

		cursor, err := DB.Collection("feeds").Find(DB.Ctx, bson.M{}, options.Find().SetSort(bson.D{{"created_at", -1}}))

		if err != nil {
			log.Fatal(err)
		}

		var feeds []models.Feed

		if err = cursor.All(DB.Ctx, &feeds); err != nil {
			log.Fatal(err)
		}

		View(w, "rss-feeds.html", feeds)

		return
	}

	View(w, "create.html", nil)
	return
}

func (fc FeedController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	validation := fc.ValidateAddFeedRequest(r)

	if len(validation) != 0 {

		json.NewEncoder(w).Encode(validation)

		return
	}

	xmlDoc, rss_link := fc.xmlDoc(r.FormValue("rss_link"))

	if xmlDoc == nil || xmlDoc.Find("rss").Length() == 0 {

		json.NewEncoder(w).Encode(map[string]string{"rss_link": "Unable to find rss feed for the provided link."})

		return
	}

	// store rss link

	Feed := models.Feed{
		Link:       rss_link,
		Created_at: time.Now(),
	}

	insertFeedResult, err := DB.Collection("feeds").InsertOne(DB.Ctx, Feed)

	if err != nil {

		json.NewEncoder(w).Encode(map[string]string{"error": "Technical Error Occured. Please try again"})

		return

	}

	// store rss feeds/posts

	var feed_posts []interface{}

	response, _ := http.Get(rss_link)

	byteValue, _ := ioutil.ReadAll(response.Body)

	rss := Rss{}

	xml.Unmarshal(byteValue, &rss)

	for _, item := range rss.Channel.Items {

		cover := "n/a"

		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(item.Description))

		if val, err := doc.Find("img").First().Attr("src"); err == true {
			cover = val
		}

		post := models.Post{
			Feed_id:     insertFeedResult.InsertedID.(primitive.ObjectID),
			Cover:       cover,
			Title:       item.Title,
			Slug:        item.Link,
			Description: item.Description,
			Author:      item.Creator,
			Pub_date:    item.PubDate,
			Created_at:  time.Now(),
		}

		feed_posts = append(feed_posts, post)

	}

	result, _ := DB.Collection("posts").InsertMany(DB.Ctx, feed_posts)

	if len(result.InsertedIDs) == 0 {

		json.NewEncoder(w).Encode(map[string]string{"error": "Technical Error Occured. Please try again"})

		DB.Collection("feeds").DeleteOne(DB.Ctx, bson.D{{"_id", insertFeedResult.InsertedID.(primitive.ObjectID)}})

		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

	return
}

func (fc FeedController) Update(w http.ResponseWriter, r *http.Request, ps router.Params) {
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

func (fc FeedController) Delete(w http.ResponseWriter, r *http.Request, ps router.Params) {
	middleware.Auth(w, r)

	// id, _ := primitive.ObjectIDFromHex(ps.ByName("id"))

	json.NewEncoder(w).Encode(map[string]bool{"success": true})

	return
}

func (fc FeedController) ValidateAddFeedRequest(r *http.Request) map[string]string {

	errors := make(map[string]string)

	if r.FormValue("rss_link") == "" {

		errors["rss_link"] = "The rss link field is required."

	} else if _, err := url.ParseRequestURI(r.FormValue("rss_link")); err != nil {

		errors["rss_link"] = "Invalid url provided."

	}

	return errors
}

func (fc FeedController) IsValidDomainName(rss_link string) bool {

	// check if rink is just domain name
	var domainRegexp = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
		]{2,3})$`)

	if strings.HasPrefix(rss_link, "http://") {
		rss_link = strings.TrimPrefix(rss_link, "http://")
	} else if strings.HasPrefix(rss_link, "https://") {
		rss_link = strings.TrimPrefix(rss_link, "https://")
	}

	if domainRegexp.MatchString(rss_link) == true {
		return true
	}

	return false
}

func (fc FeedController) xmlDoc(rss_link string) (*goquery.Document, string) {

	var (
		xmlDoc *goquery.Document
		ltf    string // link to feed
	)

	if fc.IsValidDomainName(rss_link) == true {

		possibleRssPatterns := []string{"/rss", "/feed", "/feeds/posts/default", "/rss.xml"}

		for _, val := range possibleRssPatterns {

			response, err := http.Get(rss_link + val)

			if err != nil {
				continue
			} else if response == nil {
				continue
			} else {
				defer response.Body.Close()
				//We Read the response body on the line below.
				doc, _ := goquery.NewDocumentFromReader(response.Body)
				result := doc.Find("rss").Length()

				if result >= 1 {
					xmlDoc = doc
					ltf = rss_link + val
					break
				} else {
					continue
				}

			}

		}

	} else {

		response, _ := http.Get(rss_link)

		if response != nil {

			defer response.Body.Close()
			//We Read the response body on the line below.
			doc, _ := goquery.NewDocumentFromReader(response.Body)
			result := doc.Find("rss").Length()

			if result >= 1 {
				xmlDoc = doc
				ltf = rss_link
			}
		}
	}

	return xmlDoc, ltf
}
