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
	DB, cxtCancel := config.NewDatabase(1000)

	defer cxtCancel()

	defer func() {
		if err := DB.Client().Disconnect(DB.Ctx); err != nil {
			panic(err)
		}
	}()

	// routes

	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("./public/assets"))

	fc := controllers.NewFollowingController(DB)
	lc := controllers.NewLoginController(DB)
	hc := controllers.NewHomeController(DB)

	router.GET("/", hc.Index)

	router.GET("/following", fc.Index)
	router.GET("/following/manage", fc.Create)
	router.POST("/following/manage", fc.Store)

	router.PUT("/following/:id/update", fc.Update)

	router.GET("/login", lc.Create)
	router.POST("/login", lc.Store)

	log.Fatal(http.ListenAndServe(":8080", router))

}
