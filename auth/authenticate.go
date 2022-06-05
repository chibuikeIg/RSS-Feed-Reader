package auth

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/models"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Users = map[primitive.ObjectID]models.User{}
var Sessions = map[string]primitive.ObjectID{}

func Login(user *models.User, w http.ResponseWriter) *models.User {

	// create session

	sID := uuid.NewV4()

	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}

	http.SetCookie(w, c)

	Sessions[c.Value] = user.Id
	Users[user.Id] = *user

	return user

}

func Logout(w http.ResponseWriter, r *http.Request) {

	// get cookie
	c, _ := r.Cookie("session")

	// delete session

	delete(Sessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	return

}
