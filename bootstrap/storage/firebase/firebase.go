package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func CreateAuthClient() *auth.Client {
	opt := option.WithCredentialsFile("jsons/firebase-credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("Failed to create firebase app. ERR: %s", err.Error())
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create firebase auth client. ERR: %s", err.Error())
	}

	return authClient
}

