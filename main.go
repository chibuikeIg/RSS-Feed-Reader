package main

import (
	"log"
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	quit := make(chan struct{})

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
	ffc := controllers.NewFeedController(DB, &quit)

	router.GET("/", hc.Index)

	router.GET("/following", fc.Index)
	router.PUT("/following/:id/update", fc.Update)
	router.DELETE("/following/:id/delete", fc.Delete)

	router.GET("/following/feeds", ffc.Create)
	router.POST("/following/feeds", ffc.Store)
	router.DELETE("/feeds/:feed_id/delete", ffc.Delete)

	router.GET("/login", lc.Create)
	router.POST("/login", lc.Store)
	router.GET("/logout", lc.Logout)

	// go controllers.FetchFeed(quit)

	log.Fatal(http.ListenAndServe(":8080", router))

}
