package web

import (
	"net/http"
)

// Run starts a new http server with a default Multiplexer
func Run(address string) {
	// Initialize custom logger
	app := NewApp()

	srv := &http.Server{
		Addr:     address,
		ErrorLog: app.errLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server %s", address)
	err := srv.ListenAndServe()
	app.errLog.Fatalln(err)
}
