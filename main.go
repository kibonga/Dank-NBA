package main

import (
	"Dank-NBA/api"
	"fmt"
	"sync"
)

func main() {
	players := []int{1, 2, 3}
	var wg sync.WaitGroup

	for _, id := range players {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			getPlayer(id)
		}(id)
	}

	wg.Wait()
}

func getPlayer(id int) {
	player, err := api.Get(id)
	if err != nil {
		fmt.Println("error occurred", err)
	}
	fmt.Println(player)
}
