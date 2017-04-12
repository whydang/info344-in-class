package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const defaultPort = "443"

const (
	headerContentType              = "Content-Type"
	headerAllowAccessControlOrigin = "Access-Control-Allow-Origin"
)

const (
	charsetUTF8     = "; charset=utf-8"
	contentTypeJSON = "application/json" + charsetUTF8
	contentTypeText = "text/plain" + charsetUTF8
)

type zip struct {
	Zip   string `json:"zip"`
	City  string `json:"city"`
	State string `json:"state"`
}

//zipSlice is a slice of pointers to zip structs (*zip)
type zipSlice []*zip

//zipIndex is a map of string to zipSlice
type zipIndex map[string]zipSlice

//loadZipsFromCSV loads zip records from a CSV file.
//This expects that the zip code is in position 0,
//city is in position 3, and state is in position 6.
func loadZipsFromCSV(filePath string) (zipSlice, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening zips file: %v", err)
	}

	//create a new CSV reader, which can read and parse
	//a stream of CSV data, one line at a time
	reader := csv.NewReader(f)

	//the first record is really the column names,
	//which we don't need, so just read and discard them
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV field names: %v", err)
	}

	//make a zipSlice, and preset capacity so that it
	//doesn't have to reallocate as it loads
	zips := make(zipSlice, 0, 43000)

	//read lines until we reach the end of the file
	//the .Read() method will return io.EOF when
	//it reaches the end of the file
	for {
		//read the next record
		record, err := reader.Read()
		//if we reached the end of the file,
		//return the zipSlice and no error
		if err == io.EOF {
			return zips, nil
		}
		//if we encountered some other error,
		//return it
		if err != nil {
			return nil, fmt.Errorf("error loading zips from CSV: %v", err)
		}

		//create and populate a new *zip
		//this syntax accomplishes the same thing
		//as using new() followed by field assignments
		//but does all of that in one statement
		z := &zip{
			Zip:   record[0],
			City:  record[3],
			State: record[6],
		}

		//append to the zipSlice
		zips = append(zips, z)
	}
}

//zipsForCityHandler handles request for the /zips/city/* resource
func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {
	// /zips/city/seattle
	_, city := path.Split(r.URL.Path)
	lcity := strings.ToLower(city)

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAllowAccessControlOrigin, "*")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(zi[lcity]); err != nil {
		http.Error(w, "error encoding json: "+err.Error(), http.StatusInternalServerError)
	}
}

//rootHandler handles requests for the root resource
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypeText)
	w.Write([]byte("try requesting /zips/city/seattle"))
}

//main is the entry-point for all go programs
//program execution starts with this function
func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	certPath := os.Getenv("CERTPATH")
	keyPath := os.Getenv("KEYPATH")

	zips, err := loadZipsFromCSV("zips.csv")

	//if there was an error loading the zips, report it an exit
	if err != nil {
		log.Fatal("error loading zips: " + err.Error())
	}

	fmt.Printf("loaded %d zips\n", len(zips))

	//build a map of lower-cased city name
	//to the zips in that city
	zi := make(zipIndex)
	for _, z := range zips {
		lower := strings.ToLower(z.City)
		zi[lower] = append(zi[lower], z)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/zips/city/", zi.zipsForCityHandler)

	fmt.Printf("server is listening at %s...\n", addr)

	log.Fatal(http.ListenAndServeTLS(addr, certPath, keyPath, nil))
}
