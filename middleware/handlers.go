package main

import "net/http"

const headerContentType = "Content-Type"
const contentTypeText = "text/plain"

//HelloHandler1 says hello from handler 1
func HelloHandler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypeText)
	w.Write([]byte("Hello from Handler 1"))
}

//HelloHandler2 says hello from handler 2
func HelloHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypeText)
	w.Write([]byte("Hello from Handler 2"))
}

//HelloHandler3 says hello from handler 3
func HelloHandler3(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypeText)
	w.Write([]byte("Hello from Handler 3"))
}
