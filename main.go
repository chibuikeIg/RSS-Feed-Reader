package main

import (
	"github.com/chibuikeIg/Rss_blog/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func main() {

	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("./public/assets"))
	
	uc := controllers.NewUserController()

	router.GET("/", uc.Index)

	log.Fatal(http.ListenAndServe(":8080", router))

}

