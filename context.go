package palantir

import (
	"net/http"

	"appengine/user"

	"appengine"
)

type Context struct {
	appengine.Context
	user *user.User
}

type ctxHandler func(http.ResponseWriter, *http.Request, Context) error

func handle(h ctxHandler) http.HandlerFunc {
	return _handle(h, false)
}

func handleLogged(h ctxHandler) http.HandlerFunc {
	return _handle(h, true)
}

func _handle(h ctxHandler, assertLogged bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Context{}
		c.Context = appengine.NewContext(r)

		u := user.Current(c)
		if u != nil && u.Admin {
			c.user = u
		}
		if assertLogged && c.user == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		err := h(w, r, c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
