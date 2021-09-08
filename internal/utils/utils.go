package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"
	"github.com/mmuoDev/commons/time"
	"github.com/pkg/errors"
)

var decoder = schema.NewDecoder()

// GetQueryParams maps the query params from an http request into an interface
func GetQueryParams(value interface{}, r *http.Request) error {
	// decoder lookup for values on the json tag, instead of the default schema tag
	decoder.SetAliasTag("json")

	var globalErr error

	// Decoder Register for custom type ISO8601
	decoder.RegisterConverter(time.ISO8601{}, func(input string) reflect.Value {
		ISOTime, errISO := time.NewISO8601(input)

		if errISO != nil {
			globalErr = errors.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(time.ISO8601{})
		}

		return reflect.ValueOf(ISOTime)
	})

	// Decoder Register for custom type Epoch
	decoder.RegisterConverter(time.Epoch(0), func(input string) reflect.Value {
		ISOTime, errISO := time.NewISO8601(input)

		if errISO != nil {
			globalErr = errors.Wrapf(errISO, "handler - invalid iso time provided")
			return reflect.ValueOf(time.ISO8601{}.ToEpoch())
		}

		return reflect.ValueOf(ISOTime.ToEpoch())
	})

	if err := decoder.Decode(value, r.URL.Query()); err != nil {
		return errors.Wrapf(err, "handler - failed to decode query params")
	}

	if globalErr != nil {
		return globalErr
	}

	return nil
}

//ServeError serves an error
func ServeError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

//ServeJSON serves a JSON as response
func ServeJSON(res interface{}, w http.ResponseWriter, code int) {
	bb, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bb)
}

// FileToStruct reads a json file to a struct.. a reader for the file bytes is also returned
func FileToStruct(filepath string, s interface{}) (io.Reader, error) {
	bb, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bb, s)
	return bytes.NewReader(bb), nil
}
