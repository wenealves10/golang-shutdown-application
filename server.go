package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer stop()

	var wg sync.WaitGroup

	server := &http.Server{
		Addr: ":3000",
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()

		log.Println("Closing HTTP Server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	fmt.Println("I'm leaving, Bye!")

}
