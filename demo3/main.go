/*
	Non-buffered Channel Example
*/

package main

import "fmt"

func retrieveName(c chan string) {
	var input string
	fmt.Scanln(&input)
	c <- input
}

func printName(c chan string, d chan bool) {
	name := <-c
	fmt.Println("Hello from go channels", name,
		"you may pass!")
	d <- true
}

func main() {
	ch := make(chan string)
	done := make(chan bool)

	fmt.Print("Halt! Who goes there?: ")
	go retrieveName(ch)
	go printName(ch, done)

	<-done
}
