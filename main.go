package main

import (
	"fmt"

	"github.com/openmicrotools/jwt-inspect/pkg/cmdinput"
	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	decodedJwt, err := jwt.DecodeJwt(cmdinput.HandleInput())
	if err != nil {
		fmt.Printf("An error was encountered:\n%s\n", err.Error())
	}
	fmt.Println(decodedJwt.ToString())
}
