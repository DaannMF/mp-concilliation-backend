package infrastructure

import (
	"net/http"
	"strconv"
	"strings"
)

func IsAuthorizedCallerScopes(r *http.Request) bool {
	var present bool
	scopes := getCallerScopes(r)

	for _, value := range scopes {
		present = present || value == "credits_crud"
	}
	return present
}

func GetCallerAuthorization(r *http.Request) (int64, error) {
	callerID, err := GetCallerID(r)
	if err != nil && !IsCallerAdmin(r) {
		return callerID, err
	}
	return callerID, nil
}

func IsAuthorizedClientApplication(r *http.Request, authorizedApplication string) bool {
	clientApplication := r.Header.Get("X-Api-Client-Application")
	return strings.EqualFold(clientApplication, authorizedApplication)
}

func GetCallerID(r *http.Request) (int64, error) {
	providedID := r.Header.Get("X-Caller-Id")
	if providedID == "" {
		providedID = r.URL.Query().Get("caller.id")
	}

	id, err := strconv.ParseInt(providedID, 10, 64)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func IsCallerAdmin(r *http.Request) bool {
	var present bool
	scopes := getCallerScopes(r)

	for _, value := range scopes {
		present = present || value == "admin"
	}
	return present
}

func getCallerScopes(r *http.Request) []string {
	providedCallerScopes := r.Header.Get("X-Caller-Scopes")
	if providedCallerScopes == "" {
		providedCallerScopes = r.URL.Query().Get("caller.scopes")
	}
	return strings.Split(providedCallerScopes, ",")
}
