package appengine_bootstrap

import (
	"net/http"

	"appengine/user"

	"appengine"
)

type Context struct {
	appengine.Context
	user *user.User
}

var (
	contexter = func(r *http.Request) appengine.Context {
		return appengine.NewContext(r)
	}
)

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
		c.Context = contexter(r)

		c.user = user.Current(c)
		if c.user != nil && assertLogged {
			//if c.user == nil {
			//	// No logged user, redirect to root
			//	http.Redirect(w, r, "/", http.StatusFound)
			//	return
			//}
			c.Infof("Looking for authorized account %s", c.user.Email)
			if !c.user.Admin {
				q := FindAccount(c).Filter("Email = ", c.user.Email)
				var accounts []Account
				if _, err := q.GetAll(c, &accounts); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if len(accounts) != 1 {
					http.Error(w, "Unable to find unique authorized account", http.StatusForbidden)
					return
				}
				if !accounts[0].Authorized {
					http.Error(w, "Unauthorized account", http.StatusForbidden)
					return
				}
			}
		}
		err := h(w, r, c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
