package middleware

import "firebase.google.com/go/v4/auth"

type Middleware struct {
	fireAuth *auth.Client
}

func NewMiddleware(auth *auth.Client) *Middleware {
	return &Middleware{
		fireAuth: auth,
	}
}