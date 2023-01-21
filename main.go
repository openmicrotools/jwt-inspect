package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// sample taken directly from jwt.io
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

func main() {

	args := os.Args[1:] // Args[0] is the name of the binary/program so we'll ignore it
	if len(args) != 1 {
		panic("Incorrect number of args")
	}

	jwt := args[0]

	s, _ := decodeJwt(jwt)
	fmt.Println(s)
}

func decodeJwt(s string) (string, error) {

	// header, payload, sig
	hps := strings.Split(s, ".")

	// eventual return
	var decoded string
	for _, elem := range hps {

		// decode the element into a []byte, do nothing if it blows up :)
		bytes, _ := base64.RawURLEncoding.DecodeString(elem)

		// declare some generic housing for json
		var jsonOutput map[string]interface{}

		// try to unmarshal, just move on if it blows up
		err := json.Unmarshal(bytes, &jsonOutput)
		if err != nil {
			// just skip this one, I guess it's bad
			continue
		}

		// rudamentary append onto the return
		decoded += fmt.Sprintf("%v\n", jsonOutput)

	}

	return decoded, nil

}
