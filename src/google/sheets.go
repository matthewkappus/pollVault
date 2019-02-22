package google

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)

// Service connects rosters to Google Sheets and Slides
type Service struct {
	// session store
	sheets *sheets.Service
	drive  *drive.Service
}

// ListHandler collects specifics
func (s *Service) ListHandler(w http.ResponseWriter, r *http.Request) {
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	sID := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	readRange := "Class Data!A2:E"
	resp, err := s.sheets.Spreadsheets.Values.Get(sID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Fprintf(w, "%s, %s\n", row[0], row[4])
		}
	}
}

// OA2CallbackHandler starts a Google Clasroom Service after exchanging for an oauth2 token
// It writes  a login object back to the client if using the sesin[sid] key stored on authentication
func (s *Service) OA2CallbackHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	token, err := oauth2Config.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		http.Error(w, "CallbackHandler: Token echange error error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err = s.startAPI(w, r, token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = s.setUser(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusTemporaryRedirect)
}

// an error if a Client can't be created to make service calls
func (s *Service) startAPI(w http.ResponseWriter, r *http.Request, t *oauth2.Token) error {
	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2Config.TokenSource(ctx, t))
	var err error
	s.sheets, err = sheets.New(client)
	if err != nil {
		return err
	}
	s.drive, err = drive.New(client)
	// TODO: email = get user identity aboutservice.get().do()

	return err
}

func (s *Service) setUser() error {
	if s.drive == nil {
		return fmt.Errorf("Drive service is nil")
	}
	about, err := drive.NewAboutService(s.drive).Get().Do()
	if err != nil {
		return err
	}
	log.Print("hello " + about.User.EmailAddress)
	return nil
}
