package error

import (
	"fmt"
	"testing"
)

type internalError struct {
	Name      string
	E1        error
	E2        error
	P         string
	ExpectVal error
}

func TestAppendError(t *testing.T) {
	error1 := fmt.Errorf("first error")
	error2 := fmt.Errorf("second error")
	tests := []internalError{
		{
			Name:      "Both E1 and E2 values",
			E1:        error1,
			E2:        error2,
			P:         "",
			ExpectVal: fmt.Errorf("%s; %s", error1, error2),
		},
		{
			Name:      "E1 value; no E2 value",
			E1:        error1,
			E2:        nil,
			P:         "",
			ExpectVal: error1,
		},
		{
			Name:      "E2 value; no E1 value",
			E1:        nil,
			E2:        error2,
			P:         "",
			ExpectVal: error2,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := AppendError(test.E1, test.E2)
			if fmt.Sprint(actual) != fmt.Sprint(test.ExpectVal) {
				t.Log(fmt.Sprintf("expected: \"%s\"; got: \"%s\"", test.ExpectVal, actual))
				t.FailNow()
			}
		})
	}
}

func TestPrefixError(t *testing.T) {
	error1 := fmt.Errorf("first error")
	errorNil := error(nil)
	p := "my string"
	pNil := ""
	tests := []internalError{
		{
			Name:      "Both E1 and P Values",
			E1:        error1,
			E2:        errorNil,
			P:         p,
			ExpectVal: fmt.Errorf("%s %s", p, error1.Error()),
		},
		{
			Name:      "E1 value; no P value",
			E1:        error1,
			E2:        errorNil,
			P:         pNil,
			ExpectVal: fmt.Errorf("%s %s", pNil, error1.Error()),
		},
		{
			Name:      "P value, no E1 value",
			E1:        nil,
			E2:        errorNil,
			P:         p,
			ExpectVal: errorNil,
		},
		{
			Name:      "no P value op no E1 value",
			E1:        errorNil,
			E2:        errorNil,
			P:         pNil,
			ExpectVal: errorNil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := PrefixError(test.E1, test.P)
			if fmt.Sprint(actual) != fmt.Sprint(test.ExpectVal) {
				t.Log(fmt.Sprintf("expected: \"%s\"; got: \"%s\"", test.ExpectVal, actual))
				t.FailNow()
			}
		})
	}
}
