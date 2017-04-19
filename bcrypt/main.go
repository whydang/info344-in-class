package main

import (
	"fmt"
	"os"
)

const usage = `
usage:
	bcrypt hash|verify <password> [<cost>] [<pass-hash>]

<password> is required for both 'hash' and 'verify'
<cost> is required only for 'hash'
<pass-hash> is required only for 'verify'
`

func main() {
	if len(os.Args) < 4 ||
		(os.Args[1] != "hash" && os.Args[1] != "verify") {
		fmt.Println(usage)
		os.Exit(1)
	}

}
