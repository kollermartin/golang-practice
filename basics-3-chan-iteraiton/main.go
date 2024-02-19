package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println("Time taken: ", time.Since(start))
	}()
	channel := make(chan string)

	go throwingStars(channel)
	for mess := range channel {
		fmt.Println(mess)
	}
}

func throwingStars(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	
	rounds := 3

	for i := 0; i < rounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}

	close(channel)
}