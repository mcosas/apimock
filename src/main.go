package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Load Configuration ...")
	fmt.Println("Start controllers ...")
	//	setupHandlers("users")

	fmt.Println("Starting mock on 8085...")
	http.ListenAndServe(":8085", nil)
}
