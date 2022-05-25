package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/chibuikeIg/Rss_blog/config"
)

var DB *config.Database

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

	t := template.New(filename)
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
