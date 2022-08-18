package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pushSequence() <-chan int {
	clock := time.NewTicker(2 * time.Second)

	sequencePusher := make(chan int)

	go func() {

		var i int

		for range clock.C {
			fmt.Println("I'm still running...")
			i = rand.Intn(120)
			sequencePusher <- i
		}
	}()

	return sequencePusher
}

func main() {
	sequenceChannel := pushSequence()

	for i := range sequenceChannel {
		fmt.Printf("Received %v \n", i)
		if i > 90 {
			break
		}

	}

	fmt.Println("Random Sequence Finished, Starting next Task..")

	nextTask()
}

func nextTask() {
	for {
		fmt.Println("New Task")
		time.Sleep(3 * time.Second)
	}
}
