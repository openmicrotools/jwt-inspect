package jwt

import (
	"encoding/json"
	"testing"
)

type DecodeTestCase struct {
	Name                  string
	InputJwtVal           string
	InputIsPrintEpoch     bool
	InputTimeZoneLocation string
	ExpectVal             string
	ExpectErr             bool
	ExpectErrVal          string
}

func TestDecodeJwt(t *testing.T) {

	testCases := []DecodeTestCase{
		{
			Name:                  "When input is valid token and isPrintEpoch false expecting return header and payload in json with human readable date format",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "America/Los_Angeles",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"iat":"Wed, 17 Jan 2018 17:30:22 PST","name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             false,
		},
		{
			Name:                  "When input is valid token and isPrintEpoch false expecting return header and payload with human readable date format, If no location info passed in, return UTC date format",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"iat":"Thu, 18 Jan 2018 01:30:22 UTC","name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             false,
		},
		{
			Name:                  "When input is valid token and isPrintEpoch true expecting return header and payload in json with epoch format",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
			InputIsPrintEpoch:     true,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"iat":1516239022,"name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             false,
		},
		{
			Name:                  "When input is valid token and input timezone location is UTC expecting return header and payload in json format",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
			InputIsPrintEpoch:     true,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"iat":1516239022,"name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             false,
		},
		{
			Name:                  "When input header is valid but payload is not in base64url encoded expecting return header result and error message",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.K6s7vE/2ZRUY6JQ7CbeGMn77U02AhqDd+wnK/wQ1Q9c`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"}}`,
			ExpectErr:             true,
			ExpectErrVal:          "payload section is not base64url encoded",
		},
		{
			Name:                  "When input header is valid but payload is not valid json expecting return header result and error message",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.aGVsbG8gd29ybGQ.K6s7vE/2ZRUY6JQ7CbeGMn77U02AhqDd+wnK/wQ1Q9c`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"header":{"alg":"HS256","typ":"JWT"}}`,
			ExpectErr:             true,
			ExpectErrVal:          "payload section is not valid JSON",
		},
		{
			Name:                  "When input payload is valid but header is not in base64url encoded expecting return payload result and error message",
			InputJwtVal:           `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImJlZXAiOiJib29wYm9wPz8/In0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.3FMRR0q2EpYL7JAuaYFxFz9mtjNoveYU8HtievNWsXw`,
			InputIsPrintEpoch:     true,
			InputTimeZoneLocation: "",
			ExpectVal:             `{"payload":{"iat":1516239022,"name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             true,
			ExpectErrVal:          "header section is not base64url encoded",
		},
		{
			Name:                  "When input payload is valid but header is not valid json expecting return payload result and error message",
			InputJwtVal:           `aGVsbG8gd29ybGQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.3FMRR0q2EpYL7JAuaYFxFz9mtjNoveYU8HtievNWsXw`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "America/Los_Angeles",
			ExpectVal:             `{"payload":{"iat":"Wed, 17 Jan 2018 17:30:22 PST","name":"John Doe","sub":"1234567890"}}`,
			ExpectErr:             true,
			ExpectErrVal:          "header section is not valid JSON",
		},
		{
			Name:                  "When both input payload and header are not in base64url encoded expecting return empty result and error message",
			InputJwtVal:           `eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.3FMRR0q2EpYL7JAuaYFxFz9mtjNoveYU8HtievNWsXw`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{}`,
			ExpectErr:             true,
			ExpectErrVal:          "header section is not base64url encoded; payload section is not base64url encoded",
		},
		{
			Name:                  "When both input payload and header are not valid json expecting return empty result and error message",
			InputJwtVal:           `aGVsbG8gd29ybGQ.aGVsbG8gd29ybGQ.3FMRR0q2EpYL7JAuaYFxFz9mtjNoveYU8HtievNWsXw`,
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{}`,
			ExpectErr:             true,
			ExpectErrVal:          "header section is not valid JSON; payload section is not valid JSON",
		},
		{
			Name:                  "When input is not valid in the format of a JWT expecting return empty result and error message",
			InputJwtVal:           `jwt`, // wildly invalid input
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{}`,
			ExpectErr:             true,
			ExpectErrVal:          `supplied string is not in the format of a JWT`,
		},
		{
			Name:                  "When input is empty return empty result and error message",
			InputJwtVal:           "",
			InputIsPrintEpoch:     false,
			InputTimeZoneLocation: "",
			ExpectVal:             `{}`,
			ExpectErr:             true,
			ExpectErrVal:          `supplied string is empty`,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			decoded, err := DecodeJwt(test.InputJwtVal, test.InputIsPrintEpoch, test.InputTimeZoneLocation)
			if test.ExpectErr {
				if err == nil {
					t.Logf("Test expected error but received none")
					t.Fail()
				} else if err.Error() != test.ExpectErrVal {
					t.Logf("Test failed. Expected Error:\n%s\nGot:\n%s\n", test.ExpectErrVal, err.Error())
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
				t.Logf("Test failed. Expected:\n%s\nGot:\n%s\n", test.ExpectVal, string(actual))
				t.Fail()
			}
		},
		)
	}
}
