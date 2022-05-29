module github.com/chibuikeIg/Rss_blog/controllers

go 1.18

require (
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/chibuikeIg/Rss_blog/auth v0.0.0-00010101000000-000000000000
	github.com/chibuikeIg/Rss_blog/config v0.0.0-00010101000000-000000000000
	github.com/chibuikeIg/Rss_blog/middleware v0.0.0-00010101000000-000000000000
	github.com/chibuikeIg/Rss_blog/models v0.0.0-20220521113637-c60a7fc6938e
	github.com/julienschmidt/httprouter v1.3.0
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace github.com/chibuikeIg/Rss_blog/config => ../config

replace github.com/chibuikeIg/Rss_blog/auth => ../auth

replace github.com/chibuikeIg/Rss_blog/middleware => /../middleware
