package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/impressum", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../public/templates/impressum.html")
	})
	http.HandleFunc("/datenschutzerklaerung", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../public/templates/datenschutz.html")
	})

	log.Print("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
