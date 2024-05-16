package auth

import (
	"github.com/proethics/mp-conciliation/src/api/core/entities"
	"github.com/proethics/mp-conciliation/src/api/core/entities/constants"
)

type AuthResponse struct {
	UserName string             `json:"username"`
	Role     constants.UserRole `json:"role"`
	Token    string             `json:"token"`
}

func NewAuthResponse(user entities.User, token string) AuthResponse {
	return AuthResponse{
		UserName: user.UserName,
		Role:     user.UserRole,
		Token:    token,
	}
}
