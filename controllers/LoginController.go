package controllers

import (
	"log"
	"net/http"

	"github.com/chibuikeIg/Rss_blog/auth"
	"github.com/chibuikeIg/Rss_blog/config"
	"github.com/chibuikeIg/Rss_blog/middleware"
	"github.com/chibuikeIg/Rss_blog/models"
	router "github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var errors = make(map[string]string)

type LoginController struct{}

func NewLoginController(DBConn *config.Database) *LoginController {

	DB = DBConn

	return &LoginController{}
}

func (lc LoginController) Create(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Guest(w, r)

	var validation string

	if len(errors) > 0 {

		validation = errors["credentials"]
		delete(errors, "credentials")

	}

	View(w, "login.html", map[string]string{
		"credential": validation,
	})

	return
}

func (lc LoginController) Store(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Guest(w, r)

	user, err := lc.ValidateLoginRequest(r)

	if len(err) != 0 {

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		return

	}

	auth.Login(&user, w)

	http.Redirect(w, r, "/following", http.StatusSeeOther)

	return
}

func (lc LoginController) Logout(w http.ResponseWriter, r *http.Request, _ router.Params) {

	middleware.Auth(w, r)

	auth.Logout(w, r)

	http.Redirect(w, r, "/login", http.StatusSeeOther)

	return

}

func (uc LoginController) ValidateLoginRequest(r *http.Request) (models.User, map[string]string) {

	var user models.User

	err := DB.Collection("users").FindOne(DB.Ctx, bson.D{{"email", r.FormValue("email")}}).Decode(&user)

	if err == mongo.ErrNoDocuments {

		errors["credentials"] = "The Email/Password is invalid"

	} else if err != nil {

		log.Fatalln(err)

	} else {

		err = bcrypt.CompareHashAndPassword(user.Password, []byte(r.FormValue("password")))

		if err != nil {

			errors["credentials"] = "The Email/Password is invalid"
		}

	}

	return user, errors

}
