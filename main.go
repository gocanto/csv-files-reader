package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here ....")
}
