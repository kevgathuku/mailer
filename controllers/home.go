package controller

import (
	"encoding/json"
	"github.com/kn9ts/frodo"
	"net/http"
)

// Home is an example of a controller
type Home struct {
	Frodo.Controller
}

// newsletter struct to hold parsed JSON
// It will only include exported fields in the encoded output
// and will by default use those names as the JSON keys.
// The json tags on struct fields specify key names
type newsletter struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Subscribers []string `json:"subscribers"`
}

// Index defines the Post request handler
func (h *Home) Index(w http.ResponseWriter, r *Frodo.Request) {
	// Decode the JSON from the POST request body
	decoder := json.NewDecoder(r.Body)
	// Initialize a new newsletter struct
	var t newsletter
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	w.Write([]byte(t.Title))
}
