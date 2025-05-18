package main

import (
	"log"
	"net/http"

	"github.com/donejeh/go-image-search/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/search/image", api.UploadImageHandler).Methods("POST")

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
