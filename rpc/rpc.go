package rpc

import (
	"net/http"

	"github.com/afansv/xengine/rw"
	"github.com/gorilla/mux"
)

var WriteErr = NewErrBWriter(nil)

type RPC struct {
	mux *mux.Router
}

func NewRPC() RPC {
	return RPC{mux: mux.NewRouter()}
}

type ProcedureOption struct {
	httpMethod string
}

func ProcedureOptionHTTPMethod(httpMethod string) ProcedureOption {
	return ProcedureOption{httpMethod: httpMethod}
}

func checkEngine() {
	if WriteErr == nil {
		panic("check engine: no default err writer")
	}
}

func (rp *RPC) GetHandler() http.Handler {
	checkEngine()
	return rp.mux
}

func (rp *RPC) Procedure(key string, handler rw.HandlerFunc, opts ...ProcedureOption) {
	httpMethod := http.MethodPost
	for _, opt := range opts {
		if opt.httpMethod != "" {
			httpMethod = opt.httpMethod
		}
	}

	pattern := "/" + key
	httpHandler := rw.WrapHandler(handler)

	rp.mux.HandleFunc(pattern, httpHandler).Methods(httpMethod)
}
