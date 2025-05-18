package embedding

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type EmbedRequest struct {
	ImagePath string `json:"image_path"`
}

type EmbedResponse struct {
	Embedding []float64 `json:"embedding"`
}

func GetImageEmbedding(imagePath string) ([]float64, error) {
	reqBody, _ := json.Marshal(EmbedRequest{ImagePath: imagePath})
	resp, err := http.Post("http://embedder:5000/embed", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result EmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Embedding, nil
}
