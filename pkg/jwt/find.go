package jwt

import (
	"fmt"
	"regexp"
)

const jwtRegExp = `^([a-zA-Z0-9\-_\/\+]+\.){2}([a-zA-Z0-9\-_\/\+]+)$` // This matches on base64url and base64 encoded strings so we can peel something that is close but not quite correct off the args and try to process it later for better output

// FindAndRemoveJwt will locate all JWT-like strings in the []slice and remove them. The last jwt-like string will be returned as the match if multiples are found and removed
// If no jwt-like string can be found an error is returned
func FindAndRemoveJwt(strings []string) ([]string, string, error) {
	r := regexp.MustCompile(jwtRegExp) // compile our raw string regexp

	// setup variables with scope outside the loop
	var match string
	var err error
	filtered := make([]string, 0, len(strings))

	for _, v := range strings { // range over the provided strings
		if r.MatchString(v) { // try to find something that matches a jwt
			match = v // record the match
		} else {
			filtered = append(filtered, v) // anything that didn't match get's added to our filtered list of non-jwt args
		}
	}
	if match == "" {
		err = fmt.Errorf("No valid JWT found") // if we found no matches go ahead and set our error
	}
	return filtered, match, err
}
