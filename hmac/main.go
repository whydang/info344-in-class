package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

const usage = `
usage:
	hmac sign|verify <key> <value>
`

func main() {
	if len(os.Args) < 4 ||
		(os.Args[1] != "sign" && os.Args[1] != "verify") {
		fmt.Println(usage)
		os.Exit(1)
	}
	cmd := os.Args[1]
	key := os.Args[2]
	value := os.Args[3]

	switch cmd {
	case "sign":
		v := []byte(value)
		h := hmac.New(sha256.New, []byte(key))
		h.Write(v)
		sig := h.Sum(nil)
		fmt.Println(sig)

		buff := make([]byte, len(sig)+len(v))
		copy(buff, v)
		copy(buff[len(v):], sig)
		// doesn't use / and + so url doesn't interpret it differently
		fmt.Println(base64.URLEncoding.EncodeToString(buff))
	case "verify":
		buff, err := base64.URLEncoding.DecodeString(value)
		if err != nil {
			fmt.Printf("error %v\n", err)
			os.Exit(1)
		}

		v := buff[0 : len(buff)-sha256.Size]
		sig := buff[len(buff)-sha256.Size:]

		// determine if it is valid by...?
		h := hmac.New(sha256.New, []byte(key))
		h.Write(v)
		sig2 := h.Sum(nil)

		// done in constant time so avoid giving queues about
		// how close they are to cracking sig
		if hmac.Equal(sig, sig2) {
			fmt.Println("danger danger")
		}
	}
}
