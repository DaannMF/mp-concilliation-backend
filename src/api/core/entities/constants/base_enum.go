/*
Package constants define common static values.
*/
package constants

import (
	"strings"
)

type BaseEnum interface {
	GetValues() []string
}

func GetEnumValueFromString(enum BaseEnum, value string) *string {
	for _, possibleValue := range enum.GetValues() {
		if strings.EqualFold(possibleValue, value) {
			return &possibleValue
		}
	}
	return nil
}
