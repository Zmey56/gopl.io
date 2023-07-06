//Exercise 9.5
//Write a program with two goroutines that send messages back and
//forth over two unbuffered channels in ping-pong fashion.
//How many communications per second can the program sustain?

package main

import (
	"fmt"
	"time"
)

func player(name string, in chan int, out chan int) {
	for {
		// Receive the ball
		ball := <-in

		// Increment the ball and send it back
		ball++
		out <- ball

		// Print the received ball
		fmt.Printf("%s received: %d\n", name, ball)
	}
}

func main() {
	// Create two channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Launch two goroutines
	go player("Player 1", ch1, ch2)
	go player("Player 2", ch2, ch1)

	// Start the game by sending the ball to player 1
	ch1 <- 0

	// Wait for a specified duration to measure the number of communications
	duration := 5 * time.Second
	startTime := time.Now()
	time.Sleep(duration)
	endTime := time.Now()

	// Calculate the elapsed time
	elapsedTime := endTime.Sub(startTime)

	// Calculate the number of communications per second
	communicationsPerSecond := float64(len(ch1)+len(ch2)) / elapsedTime.Seconds()

	fmt.Printf("Communications per second: %.2f\n", communicationsPerSecond)
}
