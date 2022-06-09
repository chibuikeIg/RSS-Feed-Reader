package main

import (
	"log"
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/controllers"
	"github.com/julienschmidt/httprouter"
)

func init() {
	quit := make(chan struct{})
	go controllers.FetchFeed(quit)
}

func main() {

	// DB config
	DB, cxtCancel := config.NewDatabase(1000)

	defer cxtCancel()

	defer func() {
		if err := DB.Client().Disconnect(DB.Ctx); err != nil {
			config.Log(err.Error())
		}
	}()

	// routes

	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("./public/assets"))

	fc := controllers.NewFollowingController(DB)
	lc := controllers.NewLoginController(DB)
	hc := controllers.NewHomeController(DB)
	ffc := controllers.NewFeedController(DB)
	fsc := controllers.NewFeedSettingController(DB)

	router.GET("/", hc.Index)

	router.GET("/following", fc.Index)
	router.PUT("/following/:id/update", fc.Update)
	router.DELETE("/following/:id/delete", fc.Delete)

	router.GET("/following/feeds", ffc.Create)
	router.POST("/following/feeds", ffc.Store)
	router.DELETE("/feeds/:feed_id/delete", ffc.Delete)

	router.GET("/feed/settings", fsc.Create)

	router.POST("/feed/settings", fsc.Store)

	router.GET("/login", lc.Create)
	router.POST("/login", lc.Store)
	router.GET("/logout", lc.Logout)

	log.Fatal(http.ListenAndServe(":80", router))

}
