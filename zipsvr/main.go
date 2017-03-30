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
	"path"
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

// Slices are an abstraction built on top of arrays
// See: https://blog.golang.org/go-slices-usage-and-internals
// Arrays are first class and passed by value, but slices use references to arrays

// Cross-Origin Resource Sharing (CORS)
// Origin = protocol + host + port
// browsers block cross-origin AJAX requests unless your server allows them
// Most of time we don't want random client code to access our server,
// but sometimes we want to serve to anyone as public API, like Google does.
// We can allow simple GET requests from any origin
// 	Access-Control-Allow-Origin: *
// Can constrain to specific origins if we want

// After this: guy wants us to create a client to make http calls to our server

// BUG: go path is case sensitive, and bash is not. so if your pwd is all lowercase
// but go path has capitals, need to do "cd [PATH with capitals]"

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

// *arg means receiving a pointer
// http.ResponseWriter is an interface
func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {
	// /zips/city/seattle
	_, city := path.Split(r.URL.Path)
	lcity := strings.ToLower(city)

	// Let the client know we're sending json data
	// Web by default uses UTF-8 encoding, probably fine to leave it out
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// Allow origins other than our own to make http requests to us
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Need to access database from main
	// (zi zipIndex) is a "receiver"
	// It's sort of like making this a member function of zipIndex, and defining a "this", but we call "this" "zi"
	// Basically in other languages, a reference to the object to passed in to the function as the variable "this" but it's hidden.
	// In Go and Python we need to do it explicitly.
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(zi[lcity]); err != nil {
		// We have an error, so return status code internal server error
		http.Error(w, "error encoding json: "+err.Error(), http.StatusInternalServerError)
	}
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

	// If path ends in a slash, matches any request that starts in this path
	// Needs a receiver, so we do zi.zipsForCityHandler
	http.HandleFunc("/zips/city/", zi.zipsForCityHandler)

	// prints if connection occur
	fmt.Printf("server is listening at %s...\n", addr)

	// occurs... when disconnected? meh, not sure yet. :D
	log.Fatal(http.ListenAndServe(addr, nil))
}
