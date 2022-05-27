package auth

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/models"
	uuid "github.com/satori/go.uuid"
)

var Users = map[string]models.User{}
var Sessions = map[string]string{}

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