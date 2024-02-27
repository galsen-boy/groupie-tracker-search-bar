package main

import (
	"log"
	"net/http"
	"utilities/utils"
)

func main() {
	utils.GetData()
	http.HandleFunc("/", utils.MainHandler)
	http.HandleFunc("/artist/", utils.ArtistHandler)
	http.HandleFunc("/search", utils.SearchHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("http://localhost:9000")

	http.ListenAndServe(":9000", nil)
}
