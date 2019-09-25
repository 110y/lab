package main

import (
	"context"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	metrics    = expvar.NewMap("server")
	goroutines = new(expvar.Int)
)

func main() {
	ctx := context.Background()

	metrics.Set("goroutines", goroutines)
	t := time.Tick(time.Second)
	go func() {
		for {
			select {
			case <-t:
				goroutines.Set(int64(runtime.NumGoroutine()))
			}
		}
	}()

	server := &http.Server{
		Addr: fmt.Sprintf("localhost:%d", 8080),
	}

	http.HandleFunc("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go func() {
			ticker := time.Tick(time.Second)

			for range ticker {
			}
		}()

		w.Write([]byte(fmt.Sprintf("hello")))
	}))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("failed to launch server: %s", err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	server.Shutdown(ctx)
}
