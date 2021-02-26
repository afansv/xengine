package rw

import "net/http"

type Empty struct{}

type Plugin interface {
	PreBind(i interface{}, r *http.Request) error
	PostBind(i interface{}, r *http.Request) error
}
