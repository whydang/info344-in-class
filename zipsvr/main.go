package main

// go install should turn it into a bash command?
// solution: make sure i'm in the right dir
// go install && zipsvr

// hit F12 or go to definition on menu to understand
// the methods but going to the implementation

// rare to use fixed sized arrays, but sometimes we have to
// var := [size]type{}
// slice = really small struct with a data, length, capacity
// can slice an array by var[start:end_excluded], cap determined by default reallocation
// slice declaration:
// var := []type {
//   type_val, etc
// }
// pass by value or ref if you choose, but prefer value unless need to modify

//in Go, hashtables are just maps

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// creates a structure, kind of like a class
// series of field, like C/C++
// package exporting, loading stuff from JSON to Go
// they must be exported or they cannot be seen. THUS,
// we need it to be uppercase!
// json file has lowercases, loadnig will do nothing unless..
// YOU ADD ANNOTATION to map to what you want in JSON
// dont want to export? json:"-" it i.e. storing pw but do not send it!
type zip struct {
	Zip   string `json:"zip"`
	City  string `json:"city"`
	State string `json:"state"`
}

// doesn't have to be a struct
// creates a var that is of some identity... aliasing?
type zipSlice []*zip

// declaration of map is shown below
// just pointers key to values of all zips?
type zipIndex map[string]zipSlice

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
	// require bash this command: export ADDR=localhost:8000
	if len(addr) == 0 {
		log.Fatal("please set ADDR env var")
	}

	// file stored, and err stored if any
	f, err := os.Open("../data/zips.json")

	if err != nil {
		log.Fatal("err opening zip file: " + err.Error())
	}

	// zip var allocates as array of zips?
	// capacity of 43000
	zips := make(zipSlice, 0, 43000)
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&zips); err != nil {
		log.Fatal("error decoding zips json: " + err.Error())
	}
	// if success, show number of zips in the var
	fmt.Printf("loaded %d zips\n", len(zips))

	// THIS IS THE MAP BUILDING PROCESS
	// allocate map of ptr to zips
	// TODO: gonna have to look up with make() does
	zi := make(zipIndex)

	// for each loop in Go
	// first val = index...
	// second val = value at that index
	// using _ ignores the first val
	for _, z := range zips {
		// takes string and convert to lower equiv
		lower := strings.ToLower(z.City)
		zi[lower] = append(zi[lower], z)
	}
	// test to see if map is building properly
	// city -> zips
	fmt.Printf("there are %d zips in Seattle\n", len(zi["seattle"]))

	// when someone requests /hello, handle it by calling this method
	http.HandleFunc("/hello", helloHandler)

	// prints if connection occur
	fmt.Printf("server is listening at %s...\n", addr)

	// occurs... when disconnected? meh, not sure yet. :D
	log.Fatal(http.ListenAndServe(addr, nil))
}
