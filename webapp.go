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
	r.Handle("/", handle(rootHandler)).Methods("GET")
	r.Handle("/login", handle(loginHandler)).Methods("GET")
	r.Handle("/logout", handle(logoutHandler)).Methods("GET")
	r.Handle("/register", handle(registerHandler)).Methods("GET")
	r.Handle("/expire", handle(expireHandler)).Methods("GET")
	r.Handle("/app", handleLogged(appHandler)).Methods("GET")
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
		return err
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	url, err := user.LogoutURL(c, "/")
	if err != nil {
		return err
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

// Default try duration 10j
const DefaultTryDuration = 10

func registerHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	id := r.FormValue("id")
	appName := r.FormValue("app")
	if id == "" || appName == "" {
		return missingParamError()
	}
	// Check if registration exists
	reg, err := FindRegistrationByIDAndName(c, id, appName)
	if err != nil {
		return err
	}
	if reg != nil {
		// Check if tryDuration has expired
		if reg.HasExpired() {
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
	// First, find corresponding app
	app, err := FindAppByName(c, appName)
	if err != nil {
		return err
	}
	if app == nil {
		// Not found create it
		app = &App{appName, DefaultTryDuration}
		if err = app.Save(c); err != nil {
			return err
		}
	}

	reg = &Registration{
		ID:          id,
		App:         appName,
		Date:        time.Now().Unix(),
		TryDuration: app.TryDuration,
	}
	if err = reg.Save(c); err != nil {
		return err
	}
	c.Infof("Created %+v", reg)
	fmt.Fprintf(w, "OK")
	return nil
}

func expireHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	id := r.FormValue("id")
	appName := r.FormValue("app")
	if id == "" || appName == "" {
		return missingParamError()
	}
	// Check if registration exists
	reg, err := FindRegistrationByIDAndName(c, id, appName)
	if err != nil {
		return err
	}
	if reg == nil {
		return httpError{"Registration not found", http.StatusNotFound}
	}
	reg.Date = 0
	reg.Save(c)
	fmt.Fprintf(w, "OK")
	return nil
}

func appHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	var apps []App
	if _, err := FindApp(c).GetAll(c, &apps); err != nil {
		return err
	}
	data := struct {
		User *user.User
		Apps []App
	}{c.user, apps}
	return tmpl(w, "apps.tpl", data)
}
