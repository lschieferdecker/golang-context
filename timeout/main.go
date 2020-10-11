package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Channel used to receive the result from doSomething
	ch := make(chan string)

	// Creates a new context with timeout of 3 seconds
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go doSomething(ctxTimeout, 7, ch)

	select {
	// I'm waiting to check if the context has timed out
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())

	// I'm waiting to check if I receive the results from doSomething
	case result := <-ch:
		fmt.Printf("doSomething finished: %v\n", result)
	}
}

func doSomething(ctx context.Context, timeSleep time.Duration, ch chan string) {
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * timeSleep)
	fmt.Println("Waiking up...")
	ch <- "Did something!"
}
