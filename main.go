package main

import (
	"fmt"

	"github.com/openmicrotools/jwt-inspect/internal/cmdinput"
	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

func main() {
	inputJwt, isPrintEpoch := cmdinput.HandleInput()
	//decode token passing in Local
	decodedJwt, err := jwt.DecodeJwt(inputJwt, isPrintEpoch, "Local")
	if err != nil {
		fmt.Printf("An error was encountered:\n%s\n", err.Error())
	}
	fmt.Println(jwt.ToString(decodedJwt))
}
