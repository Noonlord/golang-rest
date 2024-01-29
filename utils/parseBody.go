package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
