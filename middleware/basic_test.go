package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetJSONHeader(t *testing.T) {
	expected := "application/json; charset=utf-8"

	r, _ := http.NewRequest("GET", "", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	mock := &HTTPHandlerMock{w, r}
	handler := SetJSONHeader(mock)

	handler.ServeHTTP(mock.w, r)

	header := mock.w.Header().Get("Content-Type")
	if header != expected {
		t.Errorf("Expected %s, got `%s`", expected, header)
	}
}
