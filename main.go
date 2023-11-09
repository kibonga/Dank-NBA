package main

import (
	"Dank-NBA/api"
	"fmt"
)

func main() {
	player, err := api.Get(1)
	if err != nil {
		fmt.Println("error occurred", err)
	}
	fmt.Println(player)

}
