package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// go get -u github.com/gorilla/mux

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SomeHandler)

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("STARTED")
	log.Fatal(srv.ListenAndServe())
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
	cow := ` 
^__^
(oo)\_______
(__)\       )\
    ||----w | \
    ||     ||
`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(cow))

}
