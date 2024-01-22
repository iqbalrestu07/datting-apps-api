package domain

import "context"

type Auth struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}

type AuthUsecase interface {
	Register(ctx context.Context, auth Auth) (err error)
	Login(ctx context.Context, auth Auth) (token string, err error)
}
