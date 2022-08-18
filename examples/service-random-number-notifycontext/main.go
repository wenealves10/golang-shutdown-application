package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func pushSequence(ctx context.Context) <-chan int {
	clock := time.NewTicker(2 * time.Second)

	sequencePusher := make(chan int)

	go func() {
		var i int

		for {
			select {
			case <-clock.C:
				i = rand.Intn(120)
				sequencePusher <- i
			case <-ctx.Done():
				close(sequencePusher)
				fmt.Println("Closing the Sequence")

				if err := writeLastFromSequence(i); err != nil {
					log.Fatal(err)
				}

				return
			}
		}

	}()

	return sequencePusher
}

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sequenceChannel := pushSequence(ctx)

	for i := range sequenceChannel {
		fmt.Printf("Received %v \n", i)
	}

	fmt.Println("I am leaving, bye!")

}

func writeLastFromSequence(data int) error {
	text := fmt.Sprintf("last number for random sequence: %d", data)

	err := os.WriteFile("log.txt", []byte(text), 0644)

	if err != nil {
		return err
	}

	return nil
}
