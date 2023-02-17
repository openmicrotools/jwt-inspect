package main

import (
	"fmt"
	"os"

	"github.com/openmicrotools/jwt-inspect/pkg/cmdinput"
	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

// sample taken directly from jwt.io
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

func main() {
	decodedJwt, err := jwt.DecodeJwt(cmdinput.HandleInput(os.Args[1:]))
	if err != nil {
		fmt.Printf("An error was encountered:\n%s\n", err.Error())
	}
	fmt.Println(decodedJwt.ToString())
}
