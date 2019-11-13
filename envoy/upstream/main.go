package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := map[string]interface{}{}
		for k, v := range r.Header {
			res[strings.ToLower(k)] = v[0]
		}

		body, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"err": "failed to marshal response"}`))
		}

		w.Write(body)
	}))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "failed to launch server: %s", err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to shutdown server: %s", err.Error())
		os.Exit(1)
	}
}
