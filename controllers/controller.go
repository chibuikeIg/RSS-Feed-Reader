package controllers

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/chibuikeIg/Rss_blog/config"
	strip "github.com/grokify/html-strip-tags-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println("Feed Found")
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
