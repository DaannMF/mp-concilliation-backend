package constants

import "strings"

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

func (role UserRole) String() string {
	return strings.ToLower(string(role))
}

func (role UserRole) IsAdmin() bool {
	return role == RoleAdmin
}
