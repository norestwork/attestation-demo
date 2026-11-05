package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct{}

func NewHTTPServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run() error {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	fs := http.FileServer(http.Dir("./dist"))
	router.PathPrefix("/").Handler(fs)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://2.26.100.186:8080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
