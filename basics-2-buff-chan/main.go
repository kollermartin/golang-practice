package main

import "fmt"

func main() {
	// channel := make(chan string)

	// go func() {
	// 	channel <- "First message"
	// }()

	channel := make(chan string, 1)

	channel <- "First message"
		
	fmt.Println(<-channel) 
}