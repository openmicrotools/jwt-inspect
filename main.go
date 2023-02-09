package main

import (
	"fmt"
	"os"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

// sample taken directly from jwt.io
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

func main() {

	args := os.Args[1:] // Args[0] is the name of the binary/program so we'll ignore it
	if len(args) != 1 {
		panic("Incorrect number of args")
	}

	jwtInput := args[0]

	decodedJwt, err := jwt.DecodeJwt(jwtInput)
	if err != nil {
		fmt.Printf("An error was encountered:\n%s\n", err.Error())
	}
	fmt.Println(decodedJwt.ToString())
}
