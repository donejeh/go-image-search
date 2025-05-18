package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/donejeh/go-image-search/embedding"
	"github.com/google/uuid"
)

const UploadDir = "static/uploads"

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Limit upload size to 10MB
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid image upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure upload directory exists
	os.MkdirAll(UploadDir, os.ModePerm)

	// Generate a unique filename
	ext := filepath.Ext(header.Filename)
	filename := uuid.New().String() + ext
	fullPath := filepath.Join(UploadDir, filename)

	dst, err := os.Create(fullPath)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy uploaded file to disk
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	// Generate embedding by calling Python service
	embeddingVector, err := embedding.GetImageEmbedding(fullPath)
	if err != nil {
		http.Error(w, "Failed to get embedding", http.StatusInternalServerError)
		return
	}

	// TODO: Save embedding to Elasticsearch in the next step

	// Respond with success
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"Upload successful", "filename":"%s", "embedding": %v}`, filename, embeddingVector)
}
