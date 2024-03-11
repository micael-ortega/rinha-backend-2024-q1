package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func ValidateMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	m := strings.ToUpper(method)
	if r.Method == m {
		return true
	}

	w.Header().Add("error", fmt.Sprintf("invalid method %s", m))
	w.WriteHeader(http.StatusMethodNotAllowed)
	return false
}
