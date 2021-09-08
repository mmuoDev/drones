package app

import (
	"drones/internal/jwt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	RetrieveLocationHandler http.HandlerFunc
}

//Handler returns the main handler for this application
func (a App) Handler() http.HandlerFunc {
	router := httprouter.New()

	//books
	router.HandlerFunc(http.MethodGet, "/locations", a.RetrieveLocationHandler)

	return http.HandlerFunc(router.ServeHTTP)
}

// Options is a type for application options to modify the app
type Options func(o *OptionalArgs)

// /OptionalArgs optional arguments for this application
type OptionalArgs struct {
	GetClaims jwt.GetClaimsFunc
}

//New creates a new instance of the App
func New(options ...Options) App {
	o := OptionalArgs{
		GetClaims: jwt.GetClaims(),
	}

	for _, option := range options {
		option(&o)
	}

	retrieveLocation := RetrieveLocationHandler(o.GetClaims)

	return App{
		RetrieveLocationHandler: retrieveLocation,
	}
}
