package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	xmlDoc := fc.xmlDoc(r.FormValue("rss_link"))

	if xmlDoc == nil || xmlDoc.Find("rss").Length() == 0 {

		json.NewEncoder(w).Encode(map[string]string{"rss_link": "Unable to find rss feed for the provided link."})

		return
	}

	json.NewEncoder(w).Encode(xmlDoc.Find("rss").Length())

	return
}

func (fc FollowingController) ValidateAddFeedRequest(r *http.Request) map[string]string {

	errors := make(map[string]string)

	if r.FormValue("rss_link") == "" {

		errors["rss_link"] = "The rss link field is required."

	} else if _, err := url.ParseRequestURI(r.FormValue("rss_link")); err != nil {

		errors["rss_link"] = "Invalid url provided."

	}

	return errors
}

func (fc FollowingController) IsValidDomainName(rss_link string) bool {

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

func (fc FollowingController) xmlDoc(rss_link string) *goquery.Document {

	var xmlDoc *goquery.Document

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
			}
		}
	}

	return xmlDoc
}
