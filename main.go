package main

import (
	"github.com/kevgathuku/mailer/controllers"
	"github.com/kn9ts/frodo"
	"net/http"
)

func main() {
	// Create a new instance of Frodo
	App := Frodo.New()

	// Add the root route
	App.Get("/", func(w http.ResponseWriter, r *Frodo.Request) {
		w.Write([]byte("Hello World!!!"))
	})

	App.Post("/", &controller.Home{})

	App.Serve() // Open in browser http://localhost:3102/
}
