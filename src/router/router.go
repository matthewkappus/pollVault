package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/matthewkappus/pollVault/src/google"
	"github.com/matthewkappus/rosterUpdate/src/store"
)

// Router puts webinterface on google api services
type Router struct {
	sesh   *sessions.CookieStore
	svc    *google.Service
	roster *store.Roster
}


var cs = sessions.NewCookieStore([]byte(os.Getenv("POLLVAULT_SESSION_KEY")))

const sessionName = "PollVaultSession"

// New returns a *Router or error if sesh, roster or svc can't be initialized
func New() (*Router, error) {
	r, err := store.New("/rosters.db")
	if err != nil {
		return nil, err
	}
	return &Router{
		sesh:   cs,
		roster: r,
	}, nil
}

func (router *Router) getUser(r *http.Request) (string, error) {
	s, err := router.sesh.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	u, ok := s.Values["user"].(string)
	if !ok {
		return "", fmt.Errorf("could not find user in session")
	}
	return u, nil
}

// Class struct {
// 	ID       string  `json:"id,omitempty"`
// 	Title    string  `json:"title,omitempty"`
// 	Per      string  `json:"per,omitempty"`
// 	Teacher  string  `json:"teacher,omitempty"`
// 	Students Stu415s `json:"students,omitempty"`
// }
