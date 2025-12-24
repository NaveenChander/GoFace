package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/NaveenChander/GoFace/simulator/models"
	"github.com/gorilla/mux"
)

func StartServer() {

	r := mux.NewRouter()
	shutdownChan := make(chan struct{})
	var server *http.Server

	// Register shutdown endpoint
	r.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shutting down server..."))
		go func() {
			shutdownChan <- struct{}{}
		}()
	}).Methods("GET")
	RegisterRoutes(r)

	server = &http.Server{
		Addr:    ":" + models.EnvironmentConfig.Port,
		Handler: r,
	}

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Listening on port: " + models.EnvironmentConfig.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()

	select {
	case <-sigChan:
		fmt.Println("Received OS shutdown signal")
	case <-shutdownChan:
		fmt.Println("Received API shutdown request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}
	wg.Wait()
	fmt.Println("Server stopped gracefully")
	os.Exit(0)

}

func RegisterRoutes(r *mux.Router) {

	fmt.Println("Registering API routes")

	r.HandleFunc("/api/v1/health", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/api/v1/patrons/bulk", CreateBulkPatronHandler).Methods("POST")
	r.HandleFunc("/api/v1/dimdate/insertRange", InsertDimDateRange).Methods("POST")
}

// HealthCheckHandler responds with a simple health status.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check endpoint hit")
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(resp)
}
