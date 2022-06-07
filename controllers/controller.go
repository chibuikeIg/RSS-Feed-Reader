package controllers

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/models"
	strip "github.com/grokify/html-strip-tags-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Channel     Channel  `xml:"channel"`
	Description string   `xml:"description"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Creator     string   `xml:"dc:creator"`
}

var DB *config.Database

var fm = template.FuncMap{
	"stripTags":       stripTags,
	"TruncateByWords": TruncateByWords,
	"parseObjectId":   parseObjectId,
	"add": func(a int, b int) int {
		return a + b
	},
}

func stripTags(s string) string {
	s = strip.StripTags(s)
	return s
}

func parseObjectId(id interface{}) string {

	return id.(primitive.ObjectID).Hex()

}

func TruncateByWords(s string, maxWords int) string {
	processedWords := 0
	wordStarted := false
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRuneInString(s[i:])
		if !unicode.IsSpace(r) {
			i += width
			wordStarted = true
			continue
		}

		if !wordStarted {
			i += width
			continue
		}

		wordStarted = false
		processedWords++
		if processedWords == maxWords {
			const ending = "..."
			if (i + len(ending)) >= len(s) {
				// Source string ending is shorter than "..."
				return s
			}

			return s[:i] + ending
		}

		i += width
	}

	// Source string contains less words count than maxWords.
	return s
}

// GetAllFilePathsInDirectory : Recursively get all file paths in directory, including sub-directories.

func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// ParseDirectory : Recursively parse all files in directory, including sub-directories.
func ParseDirectory(dirpath string, filename string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}

	t := template.New(filename).Funcs(fm)

	return t.ParseFiles(paths...)
}

func View(w http.ResponseWriter, view string, data any) {

	tpl, err := ParseDirectory("./views", view)

	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	tpl.Execute(w, data)
}

func findAndStoreFeeds() {

	var settings []models.Setting

	DB.Find("settings").First(&settings)

	if len(settings) > 0 {

		t1 := settings[0].Last_poll
		t2 := time.Now()
		diff := t2.Sub(t1)
		last_poll, _ := strconv.Atoi(time.Time{}.Add(diff).Format("4"))
		polling_frequency, _ := strconv.Atoi(settings[0].Polling_frequency)
		fmt.Println("started crawling", last_poll)
		if last_poll >= polling_frequency {

			// get feeds here

			cursor, err := DB.Collection("feeds").Find(DB.Ctx, bson.M{}, options.Find().SetSort(bson.D{{"created_at", -1}}))

			if err != nil {
				log.Fatalln(err)
			}

			var feeds []models.Feed

			if err = cursor.All(DB.Ctx, &feeds); err != nil {
				log.Fatalln(err)
			}

			// get latest posts from feeds

			for _, feed := range feeds {

				var feed_posts []interface{}

				response, _ := http.Get(feed.Link)

				byteValue, _ := ioutil.ReadAll(response.Body)

				rss := Rss{}

				xml.Unmarshal(byteValue, &rss)

				for _, item := range rss.Channel.Items {

					var result bson.M
					err = DB.Collection("posts").FindOne(context.TODO(), bson.D{{"slug", item.Link}}).Decode(&result)
					if err == mongo.ErrNoDocuments {

						cover := "n/a"

						doc, _ := goquery.NewDocumentFromReader(strings.NewReader(item.Description))

						if val, err := doc.Find("img").First().Attr("src"); err == true {
							cover = val
						}

						post := models.Post{
							Feed_id:     feed.Id,
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

				}

				_, err = DB.Collection("posts").InsertMany(DB.Ctx, feed_posts)

			}

			// update last time polled

			_, err = DB.Collection("settings").UpdateOne(context.TODO(), bson.D{{"_id", settings[0].Id}}, bson.D{{"$set", models.Setting{
				Last_poll: time.Now(),
			}}})

		}

	}

	fmt.Println("background service running...")
	return
}

func FetchFeed(quit chan struct{}) {

	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			findAndStoreFeeds()
		case <-quit:
			ticker.Stop()
			return
		}
	}

}
