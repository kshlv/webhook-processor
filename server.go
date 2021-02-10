package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// A reminder for thyself
var (
	tgSupportedPorts = []int{
		80,
		88,
		443,
		8443,
	}
)

// Config ...
type Config struct {
	Addr string
}

// NewConfig ...
func NewConfig() *Config {
	addr := os.Getenv("WEBHOOK_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return &Config{
		Addr: addr,
	}
}

// Start ...
func Start(conf *Config) error {
	srv := newServer()

	return http.ListenAndServe(conf.Addr, srv)
}

type server struct {
	router       *mux.Router
	sessionStore sessions.Store
}

// newServer ...
func newServer() *server {
	return &server{
		router:       mux.NewRouter(),
		sessionStore: sessions.NewCookieStore([]byte("keypair")),
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRoutes() {
	s.router.HandleFunc("/start", s.handleStart()).Methods(http.MethodPost)
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			body, _ := json.Marshal(map[string]string{"error": "id param must be presented"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
			w.Header().Add("Content-Type", "application/json")

			return
		}

		if !searchID(id) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		log.Printf("%s\n", body)
		w.WriteHeader(http.StatusOK)

		return
	}
}

func searchID(id string) bool {
	return true
}
