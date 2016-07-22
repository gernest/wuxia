package session

import (
	"encoding/base32"
	"net/http"
	"strings"
	"time"

	"github.com/gernest/wuxia/db"
	"github.com/gernest/wuxia/models"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// QLStore is the session storage implementation for gorilla/sessions using
// embedded SQL database(ql).
type QLStore struct {
	store   *db.DB
	codecs  []securecookie.Codec
	options *sessions.Options
}

// NewQLStore initillizes QLStore with the given keyPairs
func NewQLStore(store *db.DB, path string, maxAge int, keyPairs ...[]byte) *QLStore {
	return &QLStore{
		store:  store,
		codecs: securecookie.CodecsFromPairs(keyPairs...),
		options: &sessions.Options{
			Path:   path,
			MaxAge: maxAge,
		},
	}
}

// Get fetches a session for a given name after it has been added to the registry.
func (db *QLStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(db, name)
}

// New returns a new session
func (db *QLStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(db, name)
	opts := *db.options
	session.Options = &(opts)
	session.IsNew = true

	var err error
	if c, errCookie := r.Cookie(name); errCookie == nil {
		err = securecookie.DecodeMulti(name, c.Value, &session.ID, db.codecs...)
		if err == nil {
			err = db.load(session)
			if err == nil {
				session.IsNew = false
			}
		}
	}
	return session, err
}

// Save saves the session into a ql database
func (db *QLStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	// Set delete if max-age is < 0
	if session.Options.MaxAge < 0 {
		if err := db.Delete(r, w, session); err != nil {
			return err
		}
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return nil
	}

	if session.ID == "" {
		// Generate a random session ID key suitable for storage in the DB
		session.ID = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(32)), "=")
	}

	if err := db.save(session); err != nil {
		return err
	}

	// Keep the session ID key in a cookie so it can be looked up in DB later.
	encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, db.codecs...)
	if err != nil {
		return err
	}

	http.SetCookie(w, sessions.NewCookie(session.Name(), encoded, session.Options))
	return nil
}

//load fetches a session by ID from the database and decodes its content into session.Values
func (db *QLStore) load(session *sessions.Session) error {
	s, err := models.FindSessionByKey(db.store, session.ID)
	if err != nil {
		return err
	}
	return securecookie.DecodeMulti(session.Name(), string(s.Data),
		&session.Values, db.codecs...)
}

func (db *QLStore) save(session *sessions.Session) error {
	encoded, err := securecookie.EncodeMulti(session.Name(), session.Values,
		db.codecs...)
	if err != nil {
		return err
	}
	var expiresOn time.Time
	exOn := session.Values["expires_on"]
	if exOn == nil {
		expiresOn = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
	} else {
		expiresOn = exOn.(time.Time)
		if expiresOn.Sub(time.Now().Add(time.Second*time.Duration(session.Options.MaxAge))) < 0 {
			expiresOn = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
		}
	}
	s := &models.Session{
		Key:       session.ID,
		Data:      []byte(encoded),
		ExpiresOn: expiresOn,
	}
	if session.IsNew {
		return models.CreateSession(db.store, s)
	}
	return models.UpdateSession(db.store, s.Key, s.Data)
}

func (db *QLStore) destroy(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	options := *db.options
	options.MaxAge = -1
	http.SetCookie(w, sessions.NewCookie(session.Name(), "", &options))
	for k := range session.Values {
		delete(session.Values, k)
	}
	return models.DeleteSession(db.store, session.ID)
}

// Delete deletes session.
func (db *QLStore) Delete(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return db.destroy(r, w, session)
}
