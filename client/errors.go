package client

import "fmt"

// A LWAPIError is used to identify error responses when JSON unmarshalling json from a
// byte slice.
type LWAPIError struct {
	ErrorMsg     string `json:"error,omitempty"`
	ErrorClass   string `json:"error_class,omitempty"`
	ErrorFullMsg string `json:"full_message,omitempty"`
}

// Given a LWAPIError, returns a string containing the ErrorClass and ErrorFullMsg.
func (e LWAPIError) Error() string {
	return fmt.Sprintf("%v: %v", e.ErrorClass, e.ErrorFullMsg)
}

// HasError returns boolean if ErrorClass was present or not. You can
// use this function to determine if a LWAPIRes response indicates an error or not.
func (e LWAPIError) HasError() bool {
	return e.ErrorClass != ""
}

// LWAPIRes is a convenient interface used (for example) by CallInto to ensure a passed
// struct knows how to indicate whether or not it had an error.
type LWAPIRes interface {
	Error() string
	HasError() bool
}
