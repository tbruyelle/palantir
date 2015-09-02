package palantir

import (
	"appengine/user"
	"fmt"
	"github.com/gorilla/mux"
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
	return tmpl(w, "home.tpl", data)
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

// Default try duration 15j
const DefaultTryDuration = 15

func registerHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return nil
	}
	app := r.FormValue("app")
	if app == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return nil
	}
	// Check if registration exists
	q := FindRegistration(c).Filter("ID=", id).Filter("App=", app)
	var registrations []Registration
	if _, err := q.GetAll(c, &registrations); err != nil {
		return err
	}
	if len(registrations) > 0 {
		// Registration found, find the corresponding App
		reg := registrations[0]
		q = FindApp(c).Filter("Name=", reg.App)
		var apps []App
		if _, err := q.GetAll(c, &apps); err != nil {
			return err
		}
		tryDuration := DefaultTryDuration
		if len(apps) > 1 {
			tryDuration = apps[0].TryDuration
		}
		// Check if tryDuration has expired
		if reg.HasExpired(tryDuration) {
			// tryDuration expired
			c.Infof("expired %+v", reg)
			fmt.Fprintf(w, "EXPIRED")
		} else {
			c.Infof("in try duration %+v", reg)
			fmt.Fprintf(w, "OK")
		}
		return nil
	}
	// Not found create it
	reg := &Registration{
		ID:   id,
		App:  app,
		Date: time.Now().Unix(),
	}
	err := reg.Save(c)
	if err != nil {
		return err
	}
	c.Infof("Created %+v", reg)
	fmt.Fprintf(w, "OK")
	return nil
}
