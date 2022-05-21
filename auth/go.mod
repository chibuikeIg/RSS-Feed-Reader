module github.com/chibuikeIg/Rss_blog/auth

go 1.18

replace github.com/chibuikeIg/Rss_blog/models => ../models

require (
	github.com/chibuikeIg/Rss_blog/models v0.0.0-20220521113637-c60a7fc6938e
	github.com/satori/go.uuid v1.2.0
)

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
