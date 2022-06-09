# RSS FEED READER/BLOG
#### RSS Feed Reader/Personal Blog App To Stay Up To Date With Latest Posts From Your Favorite Websites

This is a short documentation on how to set up the project on your local machine

##### Project Information/Requirement**

1. Database: MongoDB
2. Stack: Golang

##### Project Setup

1. Download/Clone the repository into your working directory
2. Create a `.env` file and follow or copy the contents of `env.example` into your `.env` file to set up your environment variables
3. Open the `main.go` file and change the `log.Fatal(http.ListenAndServe(":80", router))` to `log.Fatal(http.ListenAndServe(":8080", router))`
4. on the root directory Run `go run main.go` or `go run .` to start your server (ensure you have completed setting up your enviroment variables)
5. Open your browser and type in your project host and port number E.g `localhost:8080`.
6. Login using the testuser account you created on your mongoDB test cluser.


**To Test Live Version:** 
###### Visit
http://54.90.243.213/
Password: random
email: testuser@email.com
