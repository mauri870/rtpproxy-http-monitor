package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mauri870/rtpproxy-http-monitor/internal/rtpproxyhealth"
)

// AppHandler contains the route handlers for the application
type AppHandler struct {
	RTPProxyAddr string
}

func NewAppHandler(rtpproxyAddr string) *AppHandler {
	return &AppHandler{RTPProxyAddr: rtpproxyAddr}
}

func (app *AppHandler) Serve(addr string) error {
	r := mux.NewRouter()
	r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	// Health Check endpoint
	r.Handle("/health", app.health()).Methods("GET", "HEAD")

	return http.ListenAndServe(addr, r)
}

func (app *AppHandler) health() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := rtpproxyhealth.Check(app.RTPProxyAddr); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
	})
}
