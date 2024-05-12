package login

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/proethics/mp-conciliation/src/api/core/contracts/auth"
	"github.com/proethics/mp-conciliation/src/api/core/providers"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Execute(ctx context.Context, request auth.Request) (*string, error)
}

type Implementation struct {
	UserProvider providers.User
}

func (useCase *Implementation) Execute(ctx context.Context, request auth.Request) (*string, error) {
	user, err := useCase.UserProvider.Get(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, errors.New("error generating jwt token")
	}

	return &token, nil
}