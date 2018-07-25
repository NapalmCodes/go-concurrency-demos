/*
	Concurrency example with goroutines
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(n int) {
	start := time.Now()

	time.Sleep(time.Millisecond *
		time.Duration(rand.Intn(250))) // Simulate long operation

	end := time.Since(start)
	fmt.Println("Worker ", n, ": Finished work in ", end)
}

func main() {
	for i := 0; i < 10; i++ {
		go work(i)
	}
	var input string
	fmt.Scanln(&input)
}
