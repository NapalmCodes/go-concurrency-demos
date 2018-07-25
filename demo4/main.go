/*
	Buffered channel deadlock example (until uncommented code is active)
*/

package main

import "fmt"

func main() {
	c := make(chan int, 2) //Buffered channel
	c <- 1
	c <- 2
	c <- 3

	// c3 := func() { c <- 3 }
	// go c3()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
