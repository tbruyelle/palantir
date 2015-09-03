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
			if herr, ok := err.(httpError); ok {
				http.Error(w, herr.Error(), herr.status)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type httpError struct {
	msg    string
	status int
}

func (h httpError) Error() string {
	return h.msg
}

func missingParamError() httpError {
	return httpError{"Missing Parameter", http.StatusBadRequest}
}
