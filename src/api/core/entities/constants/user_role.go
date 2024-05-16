package constants

import "strings"

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

func (status UserRole) String() string {
	return strings.ToLower(string(status))
}
