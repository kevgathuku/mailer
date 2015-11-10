# Mailer

Mass Newsletter Mailing with Go

## Introduction


This application seeks to make sending emails in bulk easier.

It is particularly suited to sending a single email e.g. a newsletter to many users asynchronously, something Go is extremely good at.

##  Requirements

This package only requires that you have Go installed and you have configured your `$GOPATH` 

* [Installing Go](https://golang.org/doc/install)
* [Configuring $GOPATH](https://golang.org/doc/code.html)

## Installation

Assuming that your `$GOPATH` is already set, install the application by running:

`$ go install https://github.com/kevgathuku/mailer`

You should now be able to launch the app by running:

`$ mailer`

## Usage

The application by default runs as a web application on port 3102.
You can also specify the port by setting the `PORT` environment variable.

The application accepts input through a JSON POST request in the following format:

```JSON
{
	"title": "Title of the newsletter",
	"content": "HTML content of the newsletter",
	"subscribers": ["subscriber1@example.com", "subscriber1@example.com", "subscriber..n@example.com"]
}
```

Once the POST request is received, the email is sent to the subscribers specified asynchronously.

The sending progress is displayed on the command line

## Configuration

The application expects the follwing environment variables to be set:

```
SENDER_ADDRESS  # The Email to use as the FROM address
SMTP_HOST		# The SMTP host to use
SMTP_PORT		# SMTP Port to use
SMTP_USERNAME  	# SMTP username, if any
SMTP_PASSWORD  	# SMTP port, if any
```

## Contributing

I am a noob at GO and I welcome any contributions that would help make this application better.

Any help would be greatly appreciated

----

Made with &#9829; by [Kevin Ndung'u](https://github.com/kevgathuku)
