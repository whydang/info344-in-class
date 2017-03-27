package main

// go install should turn it into a bash command?
// solution: make sure that it is inside SRC folder
// go install && zipsvr

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// any type can be ref or value, ref using * like C/C++
// w = what should be responded with
// can verify if it works by go installing and running server
// then accessing via web browser
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// sends back a parsed URL stored, and then turned into an object using Query()
	// and retreive any param query string
	name := r.URL.Query().Get("name")

	// adds a header to the response
	// can be viewed via F12
	// this must come first before the body
	w.Header().Add("Content-Type", "text/plain")

	// slice of bytes vs string
	// string is a type that is a slice of bytes
	// but in png, would need a slice of byte still
	// solution: []byte("str") wraps it...and breaks it down...?
	w.Write([]byte("hello world" + name))
}

func main() {
	// set addr
	addr := os.Getenv("ADDR")
	// if not found
	if len(addr) == 0 {
		log.Fatal("please set ADDR env var")
	}

	// when someone requests /hello, handle it by calling this method
	http.HandleFunc("/hello", helloHandler)

	// prints if connection occur
	fmt.Printf("server is listening at %s...\n", addr)

	// occurs... when disconnected? meh, not sure yet. :D
	log.Fatal(http.ListenAndServe(addr, nil))
}
