package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	
	ninjas := []string{"Ninja 1", "Ninja 2", "Ninja 3"}

	for _, ninja := range ninjas {
		beeper.Add(1)
		go attack(ninja, &beeper)
	}

	beeper.Wait()
	fmt.Println("All ninjas are defeated!")

	beeper.Add(1)
	go attack("Ninja 4", &beeper)
	beeper.Wait()
	fmt.Println("extra ninja defeated!")
}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja:", evilNinja)
	beeper.Done()
}