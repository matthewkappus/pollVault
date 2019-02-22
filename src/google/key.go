package google

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var oauth2Config = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OA2_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OA2_CLIENT_SECRET"),
	// ClientID:     "419661903175-ibl4qvtb8rtkhjs30vugj8s70vgiifg6.apps.googleusercontent.com",
	// ClientSecret: "oUM46Ty8v4sfESkPheFU8RR8",
	RedirectURL: "https://matthewkappus.github.io/redirect.html",
	Scopes:      []string{sheets.SpreadsheetsScope, sheets.DriveScope},
	// Scopes:       []string{gc.ClassroomProfileEmailsScope, gc.ClassroomRostersScope, gc.ClassroomCoursesScope},
	Endpoint: google.Endpoint,
}

// Oauth2URL takes a state and returns a url to authenticate this app with Google Oauth2.
// The state value is returned from google and varified by the Callbackhandler
func Oauth2URL(state string) string {
	return oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// startSheets takes a google Oauth2Token with Classroom rights and returns
// an error if a Client can't be created to make service calls
func (svc *Service) startSheets(w http.ResponseWriter, r *http.Request, t *oauth2.Token) error {
	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2Config.TokenSource(ctx, t))
	var err error
	svc.sheets, err = sheets.New(client)
	if err != nil {
		return err
	}
	return nil
}
