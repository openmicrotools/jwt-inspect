package main

import (
	"fmt"
	"os"

	"github.com/openmicrotools/jwt-inspect/lib"
)

// sample taken directly from jwt.io
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

func main() {

	args := os.Args[1:] // Args[0] is the name of the binary/program so we'll ignore it
	if len(args) != 1 {
		panic("Incorrect number of args")
	}

	jwt := args[0]

	s, _ := lib.DecodeJwt(jwt)
	fmt.Println(s)
}
