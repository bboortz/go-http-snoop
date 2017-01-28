package main

import (
	"github.com/bboortz/go-restcache"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := restcache.NewRouter()
	router.GET("/snoop", HandlerSnoop)
	router.GET("/snoop/header", HandlerSnoopHeader)

	return router
}
