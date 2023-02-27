package util

import "testing"

type CapitalizeTestCase struct {
	Name      string
	InputVal  string
	ExpectVal string
}

func TestCapitalizeFirstChar(t *testing.T) {

	testCases := []CapitalizeTestCase{
		{
			Name:      "When input is empty expect output is empty",
			InputVal:  "",
			ExpectVal: "",
		},
		{
			Name:      "When input is a string  expect output caps the first char",
			InputVal:  "hello world",
			ExpectVal: "Hello world",
		},
		{
			Name:      "When input is a string with first as capital letter expect output is the same",
			InputVal:  "Hello world",
			ExpectVal: "Hello world",
		},
		{
			Name:      "When input is a string expect output caps the first char and not change other letters",
			InputVal:  "hello World",
			ExpectVal: "Hello World",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			actual := CapitalizeFirstChar(test.InputVal)

			if actual != test.ExpectVal {
				t.Logf("Test failed. Expected:\n%s\nGot:\n%s\n", test.ExpectVal, actual)
				t.Fail()
			}
		},
		)
	}
}
