package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
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

type jsonData map[string]interface{}

func decodeJwt(s string) (string, error) {

	// header, payload, sig
	hps := strings.Split(s, ".")

	// eventual return
	decoded := fmt.Sprintln("") // start with a newline

	for i, elem := range hps {

		var jwtSection string

		switch i {
		case 0:
			jwtSection = "Header"
		case 1:
			jwtSection = "Payload"
		default:
			jwtSection = "Unknown"
		}

		// decode the element into a []byte, do nothing if it blows up :)
		bytes, _ := base64.RawURLEncoding.DecodeString(elem)

		// declare some generic housing for json
		var unmarshalledData jsonData

		// try to unmarshal, just move on if it blows up
		err := json.Unmarshal(bytes, &unmarshalledData)
		if err != nil {
			// just skip this one, I guess it's bad
			continue
		}

		for k, v := range unmarshalledData {

			// TODO: this only works at the top level, technically a JWT could contain a nested JWT so consider handling that
			numericDate, ok := v.(float64)
			if ok { // NumericDate is the format for timestamps, golang reads it as a float64 so we can detect timestamps and format them better
				unmarshalledData[k] = time.Unix(int64(numericDate), 0).Format(time.RFC1123)
			}

		}

		formattedBytes, _ := json.MarshalIndent(unmarshalledData, "", "    ")

		decoded += fmt.Sprintf("%s:\n%s\n", jwtSection, string(formattedBytes))

	}

	return decoded, nil

}
