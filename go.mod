module github.com/chibuikeIg/Rss_blog

go 1.18

replace github.com/chibuikeIg/Rss_blog/controllers => ./controllers

require github.com/chibuikeIg/Rss_blog/controllers v0.0.0-00010101000000-000000000000

require (
	github.com/chibuikeIg/Rss_blog/auth v0.0.0-00010101000000-000000000000 // indirect
	github.com/chibuikeIg/Rss_blog/middleware v0.0.0-00010101000000-000000000000 // indirect
	github.com/chibuikeIg/Rss_blog/models v0.0.0-20220521113637-c60a7fc6938e // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
)

require (
	github.com/chibuikeIg/Rss_blog/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.9.1 // indirect
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.5 // indirect
)

replace github.com/chibuikeIg/Rss_blog/models => ./models

replace github.com/chibuikeIg/Rss_blog/config => ./config

replace github.com/chibuikeIg/Rss_blog/auth => ./auth

replace github.com/chibuikeIg/Rss_blog/middleware => ./middleware
