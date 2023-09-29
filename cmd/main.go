package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"

	"github.com/akukushkinprivate/packer/internal/handler/get_pack_sizes"
	"github.com/akukushkinprivate/packer/internal/handler/number_of_packs"
	"github.com/akukushkinprivate/packer/internal/handler/set_pack_sizes"
	"github.com/akukushkinprivate/packer/internal/packer"
)

func main() {
	packerService := packer.New()
	getPackSizesHandler := get_pack_sizes.New(packerService)
	setPackSizesHandler := set_pack_sizes.New(packerService)
	numberOfPacksHandler := number_of_packs.New(packerService)

	mux := http.NewServeMux()
	mux.Handle("/getPackSizes", getPackSizesHandler)
	mux.Handle("/setPackSizes", setPackSizesHandler)
	mux.Handle("/numberOfPacks", numberOfPacksHandler)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	// Create a new HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		log.Printf("server listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for a signal to shutdown the server
	sig := <-signalCh
	log.Printf("received signal: %v\n", sig)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v\n", err)
	}

	log.Println("server shutdown gracefully")
}
