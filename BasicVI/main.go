package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	//Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(theMine)
	//Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre"
		}
	}()
	//Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan
			fmt.Println("Frome Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted.")
		}
	}()
	<-time.After(time.Second * 2)
}
func init() {
}
