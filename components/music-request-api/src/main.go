package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := GetConfigurations()
	router := httprouter.New()

	router.POST("/submit/*query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := SubmitRequest(w, r, p, config)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	// Start the server
	http.ListenAndServe(":8080", router)
}
