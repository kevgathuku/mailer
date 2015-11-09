package main

import (
	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/kn9ts/frodo"
	"github.com/kevgathuku/mailer/controllers"
	"net/http"
)

func main() {
	// Create a new instance of Frodo
	App := Frodo.New()

	// Add the root route
	App.Get("/", func(w http.ResponseWriter, r *Frodo.Request) {
		w.Write([]byte("Hello World!!!"))
	})

	App.Post("/", &controller.Home{}, Frodo.Use{
		Method: "Index",
	})

	App.Serve() // Open in browser http://localhost:3102/
}
