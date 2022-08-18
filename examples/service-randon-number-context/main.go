package main

import (
	"context"
	"fmt"
	"math/rand"
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
				return
			}
		}

	}()

	return sequencePusher
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	sequenceChannel := pushSequence(ctx)

	for i := range sequenceChannel {
		fmt.Printf("Received %v \n", i)
		if i > 90 {
			break
		}

	}

	cancel()

	fmt.Println("Random Sequence Finished, Starting next Task..")

	nextTask()
}

func nextTask() {
	for {
		fmt.Println("New Task")
		time.Sleep(2 * time.Second)
	}
}
