package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now();
	smokeSignal := make(chan bool)

	defer func() {
		fmt.Println("Mission accomplished in ", time.Since(start))
	}()
	
	evilNinja := "Kamji"

	go attack(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	fmt.Sprintln("Ninja stars thrown at ", target);
	time.Sleep(time.Second)
	attacked <- true
}