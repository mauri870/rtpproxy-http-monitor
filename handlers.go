package main

import (
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
	r.Handle("/health", app.health()).Methods("GET")

	return http.ListenAndServe(addr, r)
}

func (app *AppHandler) health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := net.Dial("udp", app.RTPProxyAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = conn.Write([]byte("TOKEN V"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf := make([]byte, 32)
		_, err = conn.Read(buf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(string(buf))

	})
}
