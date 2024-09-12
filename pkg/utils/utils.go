package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

//utils for unmarshalling json for use

func ParseBody(r *http.Request, x interface{}) {
	//x is a struct that can be used by controllers
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
