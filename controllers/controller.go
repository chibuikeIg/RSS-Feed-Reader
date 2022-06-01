package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/chibuikeIg/Rss_blog/config"
	strip "github.com/grokify/html-strip-tags-go"
)

var DB *config.Database

var fm = template.FuncMap{
	"stripTags":       stripTags,
	"TruncateByWords": TruncateByWords,
}

func stripTags(s string) string {
	s = strip.StripTags(s)
	return s
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
