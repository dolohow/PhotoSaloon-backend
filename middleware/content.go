package middleware

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/context"

	"github.com/neo9-polska/PhotoSaloon-backend/common"
)

// BodyParser middleware parses body of request to given structure and
// sets its to context.
func BodyParser(v interface{}) Middleware {
	t := reflect.TypeOf(v)

	handler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			if err != nil {
				errMsg := common.JSONMsg(err.Error())
				http.Error(w, errMsg, http.StatusBadRequest)
				return
			}

			context.Set(r, "body", val)
			next.ServeHTTP(w, r)
		})
	}
	return handler
}
