package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/context"
)

func TestBodyParser_ValidJSON(t *testing.T) {
	expected := "{\"data\":\"testvalue\"}"

	r, _ := http.NewRequest("POST", "", strings.NewReader(expected))
	w := httptest.NewRecorder()

	mock := &HTTPHandlerMock{w, r}
	handler := BodyParser(sampleData{})(mock)

	handler.ServeHTTP(w, r)

	body := context.Get(r, "body").(*sampleData)

	if body.Data != "testvalue" {
		t.Errorf("Expected `%s`, got `%s`", expected, body.Data)
	}
}

func TestBodyParser_InvalidJSON(t *testing.T) {
	expected := "{\"data:\"invalid\"}"

	r, _ := http.NewRequest("POST", "", strings.NewReader(expected))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	mock := &HTTPHandlerMock{w, r}
	handler := BodyParser(sampleData{})(mock)

	handler.ServeHTTP(w, r)

	body := context.Get(r, "body")

	if body != nil {
		t.Errorf("Expected `%s`, to be nil", expected)
	}

	if w.Code != 400 {
		t.Errorf("Expected `%d` to be `400`", w.Code)
	}
}
