package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
	dst, err := os.Create(filepath.Join(UploadDir, filename))
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

	// Respond with success
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"Upload successful", "filename":"%s"}`, filename)
}
