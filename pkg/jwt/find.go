package jwt

import (
	"fmt"
	"regexp"
)

const base64UrlJwtRegExp = `^([a-zA-Z0-9\-_]+\.){2}([a-zA-Z0-9\-_]+)$`

func FindAndRemoveJwt(strings []string) ([]string, string, error) {
	r := regexp.MustCompile(base64UrlJwtRegExp)
	var match string
	var err error
	filtered := make([]string, 0, len(strings))
	for _, v := range strings {
		if r.MatchString(v) {
			match = v
		} else {
			filtered = append(filtered, v)
		}
	}
	if match == "" {
		err = fmt.Errorf("No valid JWT found")
	}
	return filtered, match, err
}
