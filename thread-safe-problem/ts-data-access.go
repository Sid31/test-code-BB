// Copyright (c) 2023 Sid Berraf
// Author: Sid Berraf
// Email: si.berraf@gmail.com
//
// MIT License
//

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Counter struct {
	mutex      sync.Mutex
	countValue int
}

// NewCounter initializes a new Counter and returns its pointer.
func NewCounter() *Counter {
	return &Counter{}
}

// Increment increases the counter by 1 and prints the current value.
// Parameters:
//  - goroutineID: The ID of the goroutine performing the increment.
func (c *Counter) Increment(goroutineID int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.countValue++
	fmt.Printf("Entering [Goroutine %d]: Incremented! Current value: %d\n", goroutineID, c.countValue)
	time.Sleep(10 * time.Millisecond)  // Added to slow down the process for visualization
	fmt.Printf("Exiting [Goroutine %d]: Incremented! Current value: %d\n", goroutineID, c.countValue)
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.countValue
}

// main is the entry point of the program.
// It spawns multiple goroutines to concurrently increment a shared counter and prints the final counter value.
func main() {
	const defaultGoroutines = 100

	// Fetch the desired number of goroutines from command-line argument
	var totalGoroutines int
	if len(os.Args) > 1 {
		var err error
		totalGoroutines, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid input. Using default number of goroutines:", defaultGoroutines)
			totalGoroutines = defaultGoroutines
		}
	} else {
		totalGoroutines = defaultGoroutines
	}

	counter := NewCounter()
	var wg sync.WaitGroup

	fmt.Printf("Spawning %d goroutines...\n", totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			counter.Increment(goroutineID)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nAll goroutines completed! Final Counter Value: %d\n", counter.Value())
}
