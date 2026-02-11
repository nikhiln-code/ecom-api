package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/nikhiln-code/ecom-api/internal/products"
)

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


	/** Routes */
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))	
	})

	// TODO: Initialize products service with database connection
	productsService := products.NewService()
	productsHandler := products.NewHandler(productsService)
	r.Get("/products", productsHandler.ListProducts)

	return r
}

// run 
func (app *application) run(h http.Handler) error {
	svr := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second *30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	
	log.Printf("server has started on %s", app.config.addr)
	err := http.ListenAndServe(svr.Addr, svr.Handler)
	return err;
	
}

type application struct {
	config config
	//logger
	//db driver

}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}