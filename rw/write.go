package rw

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	jData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, _ = w.Write(jData)

	return nil
}
