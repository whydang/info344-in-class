package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := "localhost:4000"

	http.HandleFunc("/v1/hello1", HelloHandler1)
	http.HandleFunc("/v1/hello2", HelloHandler2)
	http.HandleFunc("/v1/hello3", HelloHandler3)

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
