package common

import (
	"errors"
	"net/http"
)

// CheckForMissingFields checks weather body contains all given fields.
// If not it returns error and sends http response.
func CheckForMissingFields(w *http.ResponseWriter, r *http.Request,
	fields []string) error {

	for _, v := range fields {
		if r.FormValue(v) == "" {
			errorMsg := JSONMsg("Missing field '%s'", v)
			http.Error(*w, errorMsg, http.StatusBadRequest)
			return errors.New(errorMsg)
		}
	}
	return nil
}
