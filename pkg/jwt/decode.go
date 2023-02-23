package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Jwt is a struct that holds our simplified type and contains json tags so it can be marshalled to JSON
type Jwt struct {
	Header  *jsonData `json:"header,omitempty"`
	Payload *jsonData `json:"payload,omitempty"`
}

// ToString converts our Jwt type to a string or returns "" on MarshalIndent failure
func (j Jwt) ToString() string {
	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (j jsonData) ToString() string {
	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}

type jsonData map[string]interface{}

// Handle decoding a base64url encoded section of a JWT
func decodeJwtSection(s string, printEpoch bool) (*jsonData, error) {

	section := make(jsonData)

	// decode the element into a []byte, do nothing if it blows up :)
	bytes, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("section is not base64url encoded")
	}

	// try to unmarshal, just move on if it blows up
	err = json.Unmarshal(bytes, &section)
	if err != nil {
		return nil, fmt.Errorf("section is not valid JSON")
	}
	if !printEpoch { // if we are printing epoch we can skip this block, otherwise process as usual
		for k, v := range section {

			// TODO: this only works at the top level, technically a JWT could contain a nested JWT so consider handling that
			numericDate, ok := v.(float64)
			if ok { // NumericDate is the format for timestamps, golang reads it as a float64 so we can detect timestamps and format them better
				(section)[k] = time.Unix(int64(numericDate), 0).Format(time.RFC1123)
			}

		}
	}

	return &section, nil

}

// DecodeJwt accepts a string and returns our Jwt type and and error.
// This function is slightly atypical in that it may return partial Jwt data in addition to an error. This is to allow partial successes if only 1 portion of the JWT string is malformed.
func DecodeJwt(s string, printEpoch bool) (Jwt, error) {

	var jwt Jwt
	// header, payload, sig
	hps := strings.Split(s, ".")
	if len(hps) != 3 {
		// this is not a JWT
		return jwt, fmt.Errorf("supplied string is not in the format of a JWT")
	}

	var returnErr error

	for i, elem := range hps[:2] { // just ignore the sig for now

		unmarshalledData, err := decodeJwtSection(elem, printEpoch)

		switch i {
		case 0:
			jwt.Header = unmarshalledData
			returnErr = appendError(returnErr, prefixError(err, "header"))
		case 1:
			jwt.Payload = unmarshalledData
			returnErr = appendError(returnErr, prefixError(err, "payload"))
		default:
			//do nothing
		}
	}

	return jwt, returnErr

}
