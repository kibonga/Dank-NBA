package main

import (
	"Dank-NBA/api"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from a Go program\n"))
	if err != nil {
		w.Write([]byte("Hello request failed\n"))
	}
}

func getPlayer(w http.ResponseWriter, request *http.Request) {
	url := strings.Split(request.URL.String(), "/")
	id, err := strconv.Atoi(url[2])

	if err != nil {
		fmt.Println("Invalid id given")
		panic(err)
	}

	player, err := api.Get(id)
	if err != nil {
		fmt.Println("error occurred", err)
	}
	fmt.Println(player)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", hello)
	server.HandleFunc("/player/", getPlayer)

	err := http.ListenAndServe(":3000", server)
	if err != nil {
		fmt.Println("Error while opening the server")
	}
}
