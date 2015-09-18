package common

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckForMissingFields_MissingField(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) error {
		return CheckForMissingFields(&w, r, []string{"str1", "str2"})
	}

	r, _ := http.NewRequest("POST", "", strings.NewReader("str1=str1"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	err := handler(w, r)

	if err == nil {
		t.Errorf("It should return non nil error")
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Bad response error code, expected `%d`, got `%d`",
			http.StatusBadRequest, w.Code)
	}
}

func TestCheckForMissingFields_AllFieldsGiven(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) error {
		return CheckForMissingFields(&w, r, []string{"str1", "str2"})
	}

	r, _ := http.NewRequest("POST", "", strings.NewReader("str1=str1&str2=str2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	err := handler(w, r)

	if err != nil {
		t.Errorf("It should return nil error, instead got `%s`", err)
	}
}
