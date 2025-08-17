package handlers

import (
	"fmt"
	"net/http"
)

func (cfg *APIConfig) MiddlewareMetricsIncrement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (apiCfg *APIConfig) GetRequestCountHandler(w http.ResponseWriter, req *http.Request) {
	numRequests := apiCfg.fileserverHits.Load()
	responseBody := fmt.Sprintf("Hits: %d", numRequests)

	w.WriteHeader(200)
	w.Write([]byte(responseBody))
}

func (apiCfg *APIConfig) ResetRequestCountHandler(w http.ResponseWriter, req *http.Request) {
	apiCfg.fileserverHits.Store(0)

	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
