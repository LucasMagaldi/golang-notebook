package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	go contador(10)
	go contador(10)
	contador(10)
}

func contador(number int) {
	for i := range number {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
