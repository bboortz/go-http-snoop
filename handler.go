package main

import (
	"github.com/bboortz/go-restcache"
	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	//"go-rsslib"
	"fmt"
	"net/http"
	"time"
)

var headerContentTypeKey string = "Content-Type"
var headerContentTypeValue string = "application/json; charset=UTF-8"

/*
 * usage: curl http://localhost:9090/snoop
 */
func HandlerSnoop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var statusCode int = http.StatusOK

	w.WriteHeader(statusCode)
	fmt.Fprint(w, "Welcome!\n")

	restcache.LogAccess(r.Method, r.RequestURI, statusCode, start)

}

/*
 * usage: curl http://localhost:9090/snoop/header
 */
func HandlerSnoopHeader(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	start := time.Now()
	//	w.Header().Set(headerContentTypeKey, headerContentTypeValue)
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	var statusCode int = http.StatusOK

	w.WriteHeader(statusCode)
	fmt.Fprint(w, "<html><head><title>snoop</title></head><body><h1>snoop</h1><h2>Request Information</h2>\n")
	fmt.Fprint(w, "<h2>Request Information</h2><table><tr><th>KEY</th><th>VALUE</th></tr>\n")
	fmt.Fprint(w, "<tr><td>CLIENT IP:PORT</td><td>"+r.RemoteAddr+"</td></tr>\n")
	fmt.Fprint(w, "<tr><td>METHOD</td><td>"+r.Method+"</td></tr>\n")
	fmt.Fprint(w, "<tr><td>HOST</td><td>"+r.Host+"</td></tr>\n")
	fmt.Fprint(w, "<tr><td>URL</td><td>"+r.URL.String()+"</td></tr>\n")
	fmt.Fprint(w, "<tr><td>PROTOCOL</td><td>"+r.Proto+"</td></tr>\n")
	fmt.Fprint(w, "</table>\n")

	fmt.Fprint(w, "<h2>Request Header</h2><table><tr><th>KEY</th><th>VALUE</th></tr>\n")
	for k, v := range r.Header {
		fmt.Fprint(w, "<tr><td>"+k+"</td><td>"+v[0]+"</td></tr>\n")
	}
	spew.Dump(r)

	fmt.Fprint(w, "</body></html>\n")
	fmt.Fprint(w, "Welcome!\n")
	restcache.LogAccess(r.Method, r.RequestURI, statusCode, start)

}
