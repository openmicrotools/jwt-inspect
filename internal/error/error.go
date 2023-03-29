package error

import "fmt"

func AppendError(e1, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	return fmt.Errorf("%s; %s", e1.Error(), e2.Error())
}

func PrefixError(e error, p string) error {
	if e != nil {
		return fmt.Errorf("%s %s", p, e.Error())
	}
	return e
}
