package router

import (
	"net/http"

	"github.com/matthewkappus/pollVault/src/poll"
)

// PollStudentHandler gets poll_id from request and renders questions for a student (after login)
func (router *Router) PollStudentHandler(w http.ResponseWriter, r *http.Request, ballot *poll.Ballot) {

}

// PollClassHandler takes a poll_id from request and renders quests for each student in class
func (router *Router) PollClassHandler(w http.ResponseWriter, r *http.Request) {

}

// PollTeacherConsoleHandler shows student poll activity and ability to close poll
func (router *Router) PollTeacherConsoleHandler(w http.ResponseWriter, r *http.Request) {

}
