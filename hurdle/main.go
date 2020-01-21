package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server hurdle is listening on : 1402")
	fmt.Println("Hello")
	log.Fatal(http.ListenAndServe(":1401", nil))
}
