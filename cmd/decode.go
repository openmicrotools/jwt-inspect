/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type jsonData map[string]interface{}

func decodeJwt(s string, isPrettyPrint bool) (string, error) {

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

		if (isPrettyPrint) {
			for k, v := range unmarshalledData {

				// TODO: this only works at the top level, technically a JWT could contain a nested JWT so consider handling that
				numericDate, ok := v.(float64)
				if ok { // NumericDate is the format for timestamps, golang reads it as a float64 so we can detect timestamps and format them better
					unmarshalledData[k] = time.Unix(int64(numericDate), 0).Format(time.RFC1123)
				}
	
			}
		}

		formattedBytes, _ := json.MarshalIndent(unmarshalledData, "", "    ")

		decoded += fmt.Sprintf("%s:\n%s\n", jwtSection, string(formattedBytes))

	}

	return decoded, nil

}
