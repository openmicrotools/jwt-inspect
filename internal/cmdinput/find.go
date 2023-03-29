package cmdinput

import (
	"fmt"
	"regexp"
)

const jwtRegExp = `^([a-zA-Z0-9\-_\/\+]+\.){2}([a-zA-Z0-9\-_\/\+]+)$` // This matches on base64url and base64 encoded strings so we can peel something that is close but not quite correct off the args and try to process it later for better output

// FindAndRemoveJwt will locate all JWT-like strings in the []slice.
// If only a single JWT is found the match as well as a slice without the JWT is returnd.
// If multiple JWT are found then the correct behavior is undetermined so we return an error.
// If no jwt-like string can be found an error is returned
func findAndRemoveJwt(strings []string) ([]string, string, error) {
	r := regexp.MustCompile(jwtRegExp) // compile our raw string regexp

	// setup variables with scope outside the loop
	var match string
	var err error
	filtered := make([]string, 0, len(strings))

	for _, v := range strings { // range over the provided strings
		if r.MatchString(v) { // try to find something that matches a jwt
			if match != "" { // if we've already found a match the input is invalid and behavior becomes undefined
				return make([]string, 0), "", fmt.Errorf("found multiple JWT arguments, please supply only one JWT")
			}
			match = v // record the match
		} else {
			filtered = append(filtered, v) // anything that didn't match get's added to our filtered list of non-jwt args
		}
	}
	if match == "" {
		err = fmt.Errorf("no valid JWT found") // if we found no matches go ahead and set our error
	}
	return filtered, match, err
}
