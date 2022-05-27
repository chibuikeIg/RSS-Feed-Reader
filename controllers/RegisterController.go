package controllers

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	"github.com/chibuikeIg/Rss_blog/models"
	router "github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct{}

func NewRegisterController(DBConn *config.Database) *RegisterController {

	DB = DBConn

	return &RegisterController{}
}

func (rc RegisterController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Guest(w, r)

	bs, _ := bcrypt.GenerateFromPassword([]byte("random"), bcrypt.MinCost)

	user := models.User{
		Name:     "Paul Ig",
		Email:    "testuser@email.com",
		Password: bs,
	}

	_, err := DB.Collection("users").InsertOne(DB.Ctx, user)

	if err != nil {

		panic(err)

	}

	/// redirect or authenticate user

	return
}
