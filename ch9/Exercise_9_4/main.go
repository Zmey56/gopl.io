// “Exercise 9.4:
//Construct a pipeline that connects an arbitrary number of
//goroutines with channels.
//
//What is the maximum number of pipeline stages you can create without
//running out of memory?
//
//How long does a value take to transit the entire pipeline?”

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channels for each stage
	input := make(chan int)
	stage1 := make(chan int)
	stage2 := make(chan int)

	// Start the pipeline stages
	go stageOne(input, stage1)
	go stageTwo(stage1, stage2)
	go stageThree(stage2)

	// Send values into the pipeline
	for i := 0; i < 10; i++ {
		input <- i
	}

	// Close the input channel to signal the end of input
	close(input)

	// Wait for the pipeline to finish
	time.Sleep(time.Second)
}

func stageOne(input <-chan int, output chan<- int) {
	for value := range input {
		// Process value
		processedValue := value * 2

		// Pass processed value to next stage
		output <- processedValue
	}
	close(output)
}

func stageTwo(input <-chan int, output chan<- int) {
	for value := range input {
		// Process value
		processedValue := value + 1

		// Pass processed value to next stage
		output <- processedValue
	}
	close(output)
}

func stageThree(input <-chan int) {
	for value := range input {
		// Process value
		fmt.Println(value)
	}
}
