package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config := GetConfigurations()
	router := httprouter.New()

	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.POST("/submit/*query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := SubmitRequest(w, r, p, config)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	fmt.Println("Running...")
	err := http.ListenAndServe(":10010", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
