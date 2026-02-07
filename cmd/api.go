package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)
type application struct {
	config config
	//logger
	//db driver

}

//run -> graceful shiutdown -> cleanup
func (app *application) mount() http.Handler {
	r:= chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // impportant for rate limiting 
	r.Use(middleware.RealIP)// important for rate limiting and anyltics and tracing
	r.Use(middleware.Logger) // important for debugging and monitoring
	r.Use(middleware.Recoverer) //recover from the crashes


	// Set a timeout value on the request context (ctx), that will signla 
	//through ctx.Done() that the request has timmed out and further processing should be stoppped
	r.Use(middleware.Timeout(60 * time.Second))



	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))	
	})

	http.ListenAndServe(app.config.addr, r)
	return r
}

//mount
type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}