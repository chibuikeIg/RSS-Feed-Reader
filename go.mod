module github.com/chibuikeIg/Rss_blog

go 1.18

replace github.com/chibuikeIg/Rss_blog/controllers => ./controllers

require github.com/chibuikeIg/Rss_blog/controllers v0.0.0-00010101000000-000000000000

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)

replace github.com/chibuikeIg/Rss_blog/models => ./models
