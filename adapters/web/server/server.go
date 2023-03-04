package server

import (
	"github.com/codegangsta/negroni"
	"github.com/fabioods/fc_hexagonal_arch/adapters/web/handler"
	"github.com/fabioods/fc_hexagonal_arch/application"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{Service: service}
}

func (s *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(r, n, s.Service)
	http.Handle("/", r)

	server := http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
