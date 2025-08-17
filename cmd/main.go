package main

import (
	"net/http"

	"chirpy/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	apiCfg := handlers.APIConfig{}

	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))

	mux.Handle("/app/", apiCfg.MiddlewareMetricsIncrement(appHandler))
	mux.HandleFunc("/healthz", handlers.ReadinessHandler)
	mux.HandleFunc("/metrics", apiCfg.GetRequestCountHandler)
	mux.HandleFunc("/reset", apiCfg.ResetRequestCountHandler)
	server := http.Server{Handler: mux, Addr: ":8080"}
	server.ListenAndServe()
}
