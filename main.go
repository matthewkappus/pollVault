package main

import (
	"net/http"

	"github.com/matthewkappus/pollVault/src/router"
)

func main() {
	router, err := router.New()
	if err != nil {
		panic(err)
	}

	// todo: set routes in router, return mux
	http.HandleFunc("/", router.HomeHandler)
	http.HandleFunc("/oauth2callbackhandler", router.OA2CallbackHandler)
	http.HandleFunc("/list", router.ListHandler)
	http.ListenAndServe(":8080", nil)
}
