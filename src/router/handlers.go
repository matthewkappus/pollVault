package router

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/matthewkappus/pollVault/src/google"
	"github.com/matthewkappus/pollVault/src/poll"
)

var tmpl = template.Must(template.ParseGlob("tmpl/*.tmpl.html"))

// CreateBallot sends post request to service for Sheet creation or renders question form
// After a ballot is created, a webosket route receives respones by created id
func (router *Router) CreateBallot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		class, err := router.roster.SelectClassByID(r.FormValue("class_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ballot := poll.New(class).AddAnswers(r.FormValue("correct"), "incorrect1", "incorrect2")
		// todo: err := router.svc.MakePollSheet(ballot)

		router.PollBallotHandler(w, r, ballot)
	}
}

// ListHandler collects specifics
func (router *Router) ListHandler(w http.ResponseWriter, r *http.Request) {
	// u, _ := router.getUser(r)
	classes, err := router.roster.SelectClassesByTeacher("matthew.kappus@aps.edu")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "class_select.tmpl.html", classes)
}

// OA2CallbackHandler starts a Google Clasroom Service after exchanging for an oauth2 token
// It writes  a login object back to the client if using the sesin[sid] key stored on authentication
func (router *Router) OA2CallbackHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	token, err := google.Oauth2Config.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		http.Error(w, "CallbackHandler: Token echange error error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err = router.svc.StartAPI(w, r, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// todo: set user off servce getUser

	// todo: store user token

	http.Redirect(w, r, "/list", http.StatusTemporaryRedirect)
}

// HomeHandler shows classes or redirects to login
func (router *Router) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='%s'>Login</a>", google.Oauth2URL("state"))
}
