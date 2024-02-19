package main

import (
	"fmt"
)

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Naruto")
	go captainElect(ninja2, "Sasuke")

	select {
		case msg1 := <-ninja1:
			fmt.Println(msg1)
		case msg2 := <-ninja2:
			fmt.Println(msg2)
	}
	roughlyFair()
	// fmt.Println(<-ninja1)
	// fmt.Println(<-ninja2)
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}

func roughlyFair() {
	ninja1:= make(chan interface{}); close(ninja1)
	ninja2:= make(chan interface{}); close(ninja2)

	var ninja1Count, ninja2Count int
	for i := 0; i < 1000; i++ {
		select {
			case <-ninja1:
				ninja1Count++
			case <-ninja2:
				ninja2Count++
		}
	}

	fmt.Printf("Ninja 1: %d\nNinja 2: %d\n", ninja1Count, ninja2Count)
}