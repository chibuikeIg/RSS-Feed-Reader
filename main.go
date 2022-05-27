package main

import (
	"log"
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// DB config
	DB, cxtCancel := config.NewDatabase(10)

	defer cxtCancel()

	defer func() {
		if err := DB.Client().Disconnect(DB.Ctx); err != nil {
			panic(err)
		}
	}()

	// routes

	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("./public/assets"))

	hc := controllers.NewHomeController(DB)
	lc := controllers.NewLoginController(DB)

	router.GET("/", hc.Index)

	router.GET("/login", lc.Create)

	log.Fatal(http.ListenAndServe(":8080", router))

}
