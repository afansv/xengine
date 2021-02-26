package rw

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

// WrapHandler do nothing.
// This function is needed in order not to forget to return after error writing in handlers.
func WrapHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = f(w, r)
	}
}
