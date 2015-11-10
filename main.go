package main

import (
	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/joho/godotenv"
	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/kn9ts/frodo"
	"github.com/kevgathuku/mailer/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// Create a new instance of Frodo
	App := Frodo.New()

	// Add the root route
	App.Get("/", func(w http.ResponseWriter, r *Frodo.Request) {
		w.Write([]byte("Hello World!!!"))
	})

	App.Post("/", &controller.Home{}, Frodo.Use{
		Method: "Index",
	})

	// Specify PORT as an environment variable
	// Necessary for running on Heroku
	port := "3102"
	if(os.Getenv("PORT") != "") {
		port = os.Getenv("PORT")
	}
	App.ServeOnPort(port)
}
