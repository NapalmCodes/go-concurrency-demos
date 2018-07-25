/*
	Select example along with tying everything together
*/

package main

import "fmt"

func fibonacci(x, y int) (int, int) {
	x, y = y, x+y
	return x, y
}

func producer(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = fibonacci(x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func consumer(c, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go consumer(c, quit)
	producer(c, quit)
}
