package constants

import "strings"

type Resource string

const (
	Payment Resource = "PAYMENT"
	User    Resource = "PAYMENT"
)

func (status Resource) String() string {
	return strings.ToLower(string(status))
}
