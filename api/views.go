package api

import(
	// Go packages
	"encoding/json"
	"net/http"
	"fmt"
	"time"

	// Internal packages
	"github.com/sartmorgs/app-login/connectad"
)


func Access(w http.ResponseWriter, r *http.Request){
	var user User
	var current_user SessionUser
	var session_init time.Time

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Session verify
	session, err := store.Get(r, "cookie-name")
	current_user = getUser(session)

	if auth := current_user.Authenticated; !auth{
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Try auth
		connectad.SetConnectionAD(user.Username, user.Password)
		conn, err := connectad.Connect()

		if err != nil{
			fmt.Printf("Failed to connect. %s", err)
		} else{
			current_user = SessionUser{
				Username:		user.Username,
				Authenticated:	true,
			}

			session.Values["user"] = current_user

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer conn.Close()

			if err := connectad.Auth(conn); err != nil{
				fmt.Printf("%v", err)
				return
			}

			session_init = time.Now()
		}
	}

	// Access Log message
	log_session := LogAccess{
		Username: user.Username,
		RegisterDate: time.Now(),
		SessionInit: session_init,
		SessionAuthentication: current_user.Authenticated,
	}
		
	log_message, err := json.Marshal(log_session)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, "%s", log_message)
}

func Logout(w http.ResponseWriter, r *http.Request){
	var current_user SessionUser

	err := json.NewDecoder(r.Body).Decode(&current_user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Session verify
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Session Destroy
	session.Values["user"] = SessionUser{}
    session.Options.MaxAge = -1

    err = session.Save(r, w)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}