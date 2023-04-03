package main

import (
	"fmt"

	"github.com/openmicrotools/jwt-inspect/internal/cmdinput"
	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	inputJwt, isPrintEpoch := cmdinput.HandleInput()
	decodedJwt, err := jwt.DecodeJwt(inputJwt, isPrintEpoch, "")
	if err != nil {
		fmt.Printf("An error was encountered:\n%s\n", err.Error())
	}
	fmt.Println(jwt.ToString(decodedJwt))
}
