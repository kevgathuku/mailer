// Substantial portions of this file are based on postman
// which is under the MIT license.
// See https://github.com/zachlatta/postman

package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	stdMail "net/mail"
	"os"

	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/joho/godotenv"

	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/kn9ts/frodo"

	"github.com/kevgathuku/mailer/Godeps/_workspace/src/github.com/kevgathuku/postman/mail"
	"github.com/kevgathuku/mailer/Godeps/_workspace/src/gopkg.in/jordan-wright/email.v1"
)

var (
	sender, subject string
	debug           bool
	workerCount     int
)

// Home is struct holding Controller instances
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

func sendMail(from, to, subject, htmlContent string, mailer *mail.Mailer,
	debug bool, success chan *email.Email, fail chan error) {

	parsedSender, err := stdMail.ParseAddress(from)
	if err != nil {
		fail <- err
		return
	}

	parsedTo, err := stdMail.ParseAddress(to)
	if err != nil {
		fail <- err
		return
	}

	message, err := mail.NewMessage(
		parsedSender,
		parsedTo,
		subject,
		htmlContent,
	)
	if err != nil {
		fail <- err
		return
	}

	if err := mailer.Send(message); err != nil {
		fail <- err
		return
	}

	success <- message
}

// Index defines the Post request handler
func (h *Home) Index(w http.ResponseWriter, r *Frodo.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// Decode the JSON from the POST request body
	decoder := json.NewDecoder(r.Body)
	// Initialize a new newsletter struct
	var t newsletter

	decodeErr := decoder.Decode(&t)
	if decodeErr != nil {
		panic(decodeErr)
	}

	recipients := t.Subscribers
	sender := os.Getenv("SENDER_ADDRESS")
	workerCount := 8
	debug := true

	mailer := mail.NewMailer(
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	)

	jobs := make(chan string, len(recipients))
	success := make(chan *email.Email)
	fail := make(chan error)
	htmlContent := t.Content
	subject := t.Title

	// Start workers
	for i := 0; i < workerCount; i++ {
		go func() {
			for recipient := range jobs {
				sendMail(sender, recipient, subject, htmlContent, &mailer, debug, success, fail)
			}
		}()
	}

	// Send jobs to workers
	for _, recipient := range recipients {
		jobs <- recipient
	}
	close(jobs)

	for i := 0; i < len(recipients); i++ {
		select {
		case msg := <-success:
			if debug {
				fmt.Printf("\rEmailed recipient %d of %d...", i+1, len(recipients))
			} else {
				bytes, err := msg.Bytes()
				if err != nil {
					fmt.Printf("Error parsing email: %v", err)
				}
				fmt.Printf("%s\n\n\n", string(bytes))
			}
		case err := <-fail:
			fmt.Fprintln(os.Stderr, "\nError sending email:", err.Error())
			os.Exit(2)
		}
	}
	fmt.Println()

	w.Write([]byte("Newsletter Successfully Sent!"))
}
