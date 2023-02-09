package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
)

type DecodeTestInput struct {
	InputVal     string
	ExpectVal    string
	ExpectErr    bool
	ExpectErrVal string
}

func TestDecode(t *testing.T) {

	table := []DecodeTestInput{
		{
			InputVal:  `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
			ExpectVal: `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"iat":"Wed, 17 Jan 2018 20:30:22 EST","name":"John Doe","sub":"1234567890"}}`,
			ExpectErr: false,
		},
		{
			InputVal:     `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.K6s7vE/2ZRUY6JQ7CbeGMn77U02AhqDd+wnK/wQ1Q9c`,
			ExpectVal:    `{"header":{"alg":"HS256","typ":"JWT"}}`,
			ExpectErr:    true,
			ExpectErrVal: "payload section is not base64url encoded",
		},
		{
			InputVal:     `jwt`, // wildly invalid input
			ExpectVal:    `{}`,
			ExpectErr:    true,
			ExpectErrVal: `supplied string is not in the format of a JWT`,
		},
	}

	for _, test := range table {
		decoded, err := DecodeJwt(test.InputVal)
		if test.ExpectErr {
			if err == nil {
				fmt.Println("Test expected error but recieved none")
				t.Fail()
			} else if err.Error() != test.ExpectErrVal {
				fmt.Printf("Test failed. Expected Error:\n%s\nGot:\n%s\n", test.ExpectErrVal, err.Error())
				t.Fail()
			}
		} else {
			if err != nil {
				t.Fail()
			}
		}
		actual, err := json.Marshal(decoded)
		if err != nil {
			t.Fail()
		} else if string(actual) != test.ExpectVal {
			fmt.Printf("Test failed. Expected:\n%s\nGot:\n%s\n", test.ExpectVal, string(actual))
			t.Fail()
		}
	}
}
