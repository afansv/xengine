package rw

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type RequestWithPlugins interface {
	Plugins() []Plugin
}

// Bind binds request body and query params to custom struct and then validates it.
func Bind(i interface{}, r *http.Request) error {
	t := reflect.TypeOf(i)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	plugs := make([]Plugin, 0)

	hasPlugins := t.Implements(reflect.TypeOf((*RequestWithPlugins)(nil)).Elem())
	if hasPlugins {
		for _, plug := range i.(RequestWithPlugins).Plugins() {
			plugs = append(plugs, plug)
		}
	}

	for _, plug := range plugs {
		if err := plug.PreBind(i, r); err != nil {
			return err
		}
	}

	if len(b) != 0 {
		if err := json.Unmarshal(b, i); err != nil {
			return err
		}
	}

	if err := decoder.Decode(i, r.PostForm); err != nil {
		return err
	}

	if err := decoder.Decode(i, r.URL.Query()); err != nil {
		return err
	}

	for _, plug := range plugs {
		if err := plug.PostBind(i, r); err != nil {
			return err
		}
	}

	return nil
}
