package main

import (
	"fmt"
	"net/http"

	"github.com/matthewkappus/pollVault/src/google"
	"github.com/matthewkappus/pollVault/src/router"
)

func main() {
	router, err := router.New()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", HomeHandler)
	// http.HandleFunc("/oauth2callbackhandler", s.OA2CallbackHandler)
	http.HandleFunc("/list", router.ListHandler)
	http.ListenAndServe(":8080", nil)
}

// HomeHandler shows classes or redirects to login
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='%s'>Login</a>", google.Oauth2URL("state"))
}
