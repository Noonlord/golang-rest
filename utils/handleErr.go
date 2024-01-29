package utils

import (
	"fmt"
	"net/http"
)

func HandleErr(err error, statusCode int, w http.ResponseWriter) bool {
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(statusCode), statusCode)
		return true
	}
	return false
}
