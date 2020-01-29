package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/110y/lab/envoy/authority/handler"
)

var rsaPrivateKey *rsa.PrivateKey

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()

	mux.Handle("/jwt", &handler.JWT{PrivateKey: rsaPrivateKey})
	mux.Handle("/jwks", &handler.JWKS{PrivateKey: rsaPrivateKey})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("failed to launch server: %s", err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown server: %s", err.Error())
		os.Exit(1)
	}
}

func init() {
	k, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to launch server: %s", err.Error())
		os.Exit(1)
	}
	rsaPrivateKey = k
}
