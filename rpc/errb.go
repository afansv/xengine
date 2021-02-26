package rpc

import (
	"net/http"
	"strings"

	"github.com/afansv/xengine/errb"
	"github.com/afansv/xengine/rw"
)

type ErrorWriter func(w http.ResponseWriter, err error) error

type Error struct {
	Code    string      `json:"code"`
	Title   string      `json:"title"`
	Details interface{} `json:"details"`
}

type ErrBWriterOpts struct {
	debug                 bool
	defaultCode           string
	defaultTitle          string
	defaultHTTPStatusCode int
	lang                  Lang
}

func ErrBWriterOptionDebug(debug bool) ErrBWriterOpts {
	return ErrBWriterOpts{
		debug: debug,
	}
}

func ErrBWriterOptionDefaultCode(defaultCode string) ErrBWriterOpts {
	return ErrBWriterOpts{
		defaultCode: defaultCode,
	}
}

func ErrBWriterOptionDefaultTitle(defaultTitle string) ErrBWriterOpts {
	return ErrBWriterOpts{
		defaultTitle: defaultTitle,
	}
}

func ErrBWriterOptionDefaultHTTPStatusCode(defaultHTTPStatusCode int) ErrBWriterOpts {
	return ErrBWriterOpts{
		defaultHTTPStatusCode: defaultHTTPStatusCode,
	}
}

func ErrBWriterOptionLang(lang Lang) ErrBWriterOpts {
	return ErrBWriterOpts{
		lang: lang,
	}
}

type StatusResolver func(code string) int

func NewErrBWriter(statRes StatusResolver, opts ...ErrBWriterOpts) ErrorWriter {
	// Final opts.
	fOpt := ErrBWriterOpts{
		defaultCode:           "unexpected",
		defaultTitle:          "unexpected error",
		defaultHTTPStatusCode: http.StatusInternalServerError,
		lang:                  LangEN,
	}

	for _, opt := range opts {
		switch {
		case opt.debug:
			fOpt.debug = true
		case opt.defaultCode != "":
			fOpt.defaultCode = opt.defaultCode
		case opt.defaultTitle != "":
			fOpt.defaultTitle = opt.defaultTitle
		case opt.defaultHTTPStatusCode != 0:
			fOpt.defaultHTTPStatusCode = opt.defaultHTTPStatusCode
		case opt.lang != "":
			fOpt.lang = opt.lang
		}
	}

	return func(w http.ResponseWriter, err error) error {
		httpStatus := http.StatusInternalServerError

		resp := Error{
			Code:  fOpt.defaultCode,
			Title: fOpt.defaultTitle,
		}
		if fOpt.debug {
			resp.Details = err.Error()
		}

		eb, ok := err.(errb.ErrBin)
		if ok {
			resp = Error{Code: eb.Code(), Details: eb.Details}

			if statRes != nil {
				httpStatus = statRes(eb.Code())
			}

			switch fOpt.lang {
			case LangRU:
				resp.Title = eb.ErrorRU()
			default:
				resp.Title = eb.Error()
			}
		}

		resp.Title = strings.Title(resp.Title)
		if !strings.HasSuffix(resp.Title, ".") {
			resp.Title += "."
		}

		w.WriteHeader(httpStatus)
		return rw.WriteResponse(w, resp)
	}
}
