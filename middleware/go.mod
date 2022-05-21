module github.com/chibuikeIg/Rss_blog/middleware

go 1.18

replace github.com/chibuikeIg/Rss_blog/auth => ../auth

require github.com/chibuikeIg/Rss_blog/auth v0.0.0-00010101000000-000000000000

require (
	github.com/chibuikeIg/Rss_blog/models v0.0.0-20220521113637-c60a7fc6938e // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
)
