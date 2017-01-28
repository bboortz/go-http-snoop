package restcache

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	//	log := logging.MustGetLogger("go-router")
	router := httprouter.New()
	router.GET("/", HandlerIndexRead)
	router.GET("/alive", HandlerAliveRead)
	router.NotFound = NotFoundHandler()
	router.MethodNotAllowed = MethodNotAllowedHandler()

	return router
}
