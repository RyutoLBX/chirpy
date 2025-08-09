package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	mux := http.NewServeMux()
	apiCfg := apiConfig{}

	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(appHandler))
	mux.HandleFunc("/healthz", readinessHandler)
	mux.HandleFunc("/metrics", apiCfg.getRequestCountHandler)
	mux.HandleFunc("/reset", apiCfg.resetRequestCountHandler)
	server := http.Server{Handler: mux, Addr: ":8080"}
	server.ListenAndServe()
}

func readinessHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (apiCfg *apiConfig) getRequestCountHandler(w http.ResponseWriter, req *http.Request) {
	numRequests := apiCfg.fileserverHits.Load()
	responseBody := fmt.Sprintf("Hits: %d", numRequests)

	w.WriteHeader(200)
	w.Write([]byte(responseBody))
}

func (apiCfg *apiConfig) resetRequestCountHandler(w http.ResponseWriter, req *http.Request) {
	apiCfg.fileserverHits.Store(0)

	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
