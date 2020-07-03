package api

import(
	"encoding/gob"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	// store will hold all session data
    store *sessions.CookieStore
)

// User struct to auth
type User struct {
	Username      	string
	Password		string
}

// User session config
type SessionUser struct {
	Username      string
	Authenticated bool
}

// Access Log struct
type LogAccess struct{
	Username				string
	RegisterDate			time.Time
	SessionInit				time.Time
	SessionEnd				time.Time
	SessionAuthentication	bool
}


func init(){
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		Path: 		"/",
		MaxAge: 	60 * 15,
		HttpOnly: 	true,
	}

	gob.Register(SessionUser{})
}

func getUser(s *sessions.Session) SessionUser {
	val := s.Values["user"]
	var user = SessionUser{}
	user, ok := val.(SessionUser)
	if !ok {
		return SessionUser{Authenticated: false}
	}
	return user
}