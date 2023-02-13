package jwt

import "fmt"

func appendError(e1, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	return fmt.Errorf("%s; %s", e1.Error(), e2.Error())
}

func prefixError(e error, p string) error {
	if e != nil {
		return fmt.Errorf("%s %s", p, e.Error())
	}
	return e
}
