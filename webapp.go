package appengine_bootstrap

import (
	"appengine/user"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

func init() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.Handle("/", handleLogged(rootHandler)).Methods("GET")
	r.Handle("/login", handle(loginHandler)).Methods("GET")
	r.Handle("/logout", handle(logoutHandler)).Methods("GET")
	r.Handle("/register", handle(registerHandler)).Methods("GET")
}

func rootHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	if r.URL.Path != "/" {
		c.Errorf("Unknow root %s", r.URL.Path)
		http.NotFound(w, r)
		return nil
	}
	tmpl, err := template.ParseFiles("templates/root.tpl", "templates/home.tpl")
	if err != nil {
		return err
	}

	registrations := make([]Registration, 0, 20)
	if c.user != nil {
		// Fetch registrations
		q := FindRegistration(c).Order("-Date").Limit(20)
		if _, err := q.GetAll(c, &registrations); err != nil {
			return err
		}
	}

	data := struct {
		User          *user.User
		Registrations []Registration
	}{
		c.user,
		registrations,
	}
	tmpl.Execute(w, data)
	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	url, err := user.LoginURL(c, "/")
	if err != nil {
		return nil
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	url, err := user.LogoutURL(c, "/")
	if err != nil {
		return nil
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	id := r.FormValue("ID")
	if id == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return nil
	}
	reg := &Registration{
		ID:      id,
		Account: c.user.Email,
		Date:    time.Now().Unix(),
	}
	reg.Save(c)
	return nil
}
