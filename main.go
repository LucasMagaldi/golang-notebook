package main

import (
	"fmt"
	"net/http"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("The Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	canal := make(chan int)
	workersQuantity := 10

	for i := range workersQuantity {
		go worker(i, canal)
	}

	for i := range 110 {
		canal <- i
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
