/*
Package errors implements commons use case error types.
*/
package errors

import (
	"errors"
	"fmt"
	"sort"

	"github.com/go-playground/validator/v10"
)

type (
	Message    string
	Parameters map[string]interface{}
)

const (
	// Error
	ErrorRunningApplication       Message = "Error running application."
	ErrorDataBaseConnection       Message = "Error trying to connect to database."
	ErrorDataBaseMigration        Message = "Error trying to create migration to database."
	ErrorCreatingConfigClient     Message = "Error trying to create configuration client."
	ErrorBindingRequest           Message = "Error binding request."
	ErrorInvalidID                Message = "Property payment_id must be a valid number."
	ErrorBeginTransaction         Message = "Error trying to begin transaction."
	ErrorCommitTransaction        Message = "Error trying to commit transaction."
	ErrorRecoverFunction          Message = "An error occurred that involved calling the recover function and rollback transaction."
	ErrorGettingResource          Message = "Error getting resource."
	ErrorCreatingResource         Message = "Error creating resource."
	ErrorUpdatingResource         Message = "Error updating resource."
	ErrorPaymentAlreadyConcillied Message = "Error payment already concillied."
	ErrorDeletingResource         Message = "Error deleting resource."
	ErrorRecordNotFound           Message = "Error record not found."
	ErrorUserNotPresent           Message = "Error user not present in context."

	// Info
	InfoDataBaseConnection Message = "Connection to the database, was successful."
)

func (message Message) GetMessage() string {
	return string(message)
}

func (message Message) GetMessageWithParams(params Parameters) string {
	msg := message.GetMessage()
	keys := make([]string, 0)

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		m := fmt.Sprintf(" %v:%v", k, params[k])
		msg += m
	}

	return msg
}

func GetMessageParamsFromValidationErrors(err error) Parameters {
	var errs validator.ValidationErrors
	params := Parameters{}

	if errors.As(err, &errs) {
		for _, f := range errs {
			params[f.Field()] = fmt.Sprintf("'%v'", f.Value())
		}
	}

	return params
}
