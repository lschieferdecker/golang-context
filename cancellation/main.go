package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cancel the context after two seconds
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	go doSomething(ctx, 3)
	go iterateOverSomething(ctx, 8)

	fmt.Scanln()
}

func doSomething(ctx context.Context, timeSleep time.Duration) {
	fmt.Println("Starting doSomething")
	select {
	// Here I'm simulating the time doSomething will take to process.
	case <-time.After(time.Second * timeSleep):
		fmt.Println("Finished doSomething")
	// Check the context Done to get cancellation signal
	case <-ctx.Done():
		fmt.Println("Cancelling doSomething")
	}
}

func iterateOverSomething(ctx context.Context, iterations int) {
	for i := 0; i < iterations; i++ {
		if ctx.Err() != nil {
			fmt.Println("Leaving iterateOverSomething...")
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println("Iteration ", i)
	}
}
