package restcache

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	//	"github.com/davecgh/go-spew/spew"
)

var headerContentTypeKey string = "Content-Type"
var headerContentTypeValue string = "application/json; charset=UTF-8"

/*
 * usage: curl -H "Content-Type: application/json" http://localhost:9090
 */
func HandlerIndexRead(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()
	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	var statusCode int = http.StatusOK
	var result Api

	result = Api{ApiName: "go-router", ApiVersion: "0.1"}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
	LogAccess(r.Method, r.RequestURI, statusCode, start)
}

/*
 * usage: curl -H "Content-Type: application/json" http://localhost:9090
 */
func HandlerAliveRead(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()
	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	var statusCode int = http.StatusOK
	var result Alive

	result = Alive{Alive: true}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
	LogAccess(r.Method, r.RequestURI, statusCode, start)
}

/*
 * error handler
 */

func NotFound(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	var statusCode int = http.StatusNotFound
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(JsonErr{Code: statusCode, Text: "Not Found"}); err != nil {
		panic(err)
	}
	LogAccess(r.Method, r.RequestURI, statusCode, start)
}
func NotFoundHandler() http.Handler { return http.HandlerFunc(NotFound) }

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	var statusCode int = http.StatusMethodNotAllowed
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(JsonErr{Code: statusCode, Text: "Method Not Allowed"}); err != nil {
		panic(err)
	}
	LogAccess(r.Method, r.RequestURI, statusCode, start)
}
func MethodNotAllowedHandler() http.Handler { return http.HandlerFunc(MethodNotAllowed) }
