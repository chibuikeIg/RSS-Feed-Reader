package middleware

import (
	"net/http"

	"github.com/chibuikeIg/Rss_blog/auth"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")

	if err != nil {

		http.Redirect(w, r, "/login", http.StatusSeeOther)

		return
	}

	user_id := auth.Sessions[c.Value]

	_, ok := auth.Users[user_id]

	if !ok {

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	return
}

func Guest(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")

	if err != nil {
		return
	}

	user_id := auth.Sessions[c.Value]

	_, ok := auth.Users[user_id]

	if ok {

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	return

}
