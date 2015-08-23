package appengine_bootstrap

import (
	"appengine"
	"appengine/datastore"
)

type App struct {
	Name        string
	TryDuration int64
}

func appKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "App", "default_app", 0, nil)
}

func (a *App) Save(c appengine.Context) error {
	key := datastore.NewIncompleteKey(c, "App", appKey(c))
	_, err := datastore.Put(c, key, a)
	return err
}

func FindApp(c appengine.Context) *datastore.Query {
	return datastore.NewQuery("App").Ancestor(appKey(c))
}

type Account struct {
	Email      string
	Authorized bool
}

func accountKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Account", "default_account", 0, nil)
}

func (a *Account) Save(c appengine.Context) error {
	key := datastore.NewIncompleteKey(c, "Account", accountKey(c))
	_, err := datastore.Put(c, key, a)
	return err
}

func FindAccount(c appengine.Context) *datastore.Query {
	return datastore.NewQuery("Account").Ancestor(accountKey(c))
}

type Registration struct {
	ID   string
	App  string
	Date int64
}

func registrationKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Registration", "default_registration", 0, nil)
}

func (r *Registration) Save(c appengine.Context) error {
	key := datastore.NewIncompleteKey(c, "Registration", registrationKey(c))
	_, err := datastore.Put(c, key, r)
	return err
}

func FindRegistration(c appengine.Context) *datastore.Query {
	return datastore.NewQuery("Registration").Ancestor(registrationKey(c))
}
