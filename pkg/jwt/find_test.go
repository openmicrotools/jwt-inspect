package jwt_test

import (
	"testing"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
)

const testJwt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
const testBadJwt = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lPz8/IiwiaWF0IjoxNTE2MjM5MDIyfQ.K6s7vE/2ZRUY6JQ7CbeGMn77U02AhqDd+wnK/wQ1Q9c`

type FindAndRemoveJwtTestCase struct {
	Name        string
	Input       []string
	ExpectSlice []string
	ExpectMatch string
	ExpectErr   bool
	ErrVal      string
}

func TestFindAndRemoveJwt(t *testing.T) {

	testcases := []FindAndRemoveJwtTestCase{
		{
			Name:        "When a single JWT is present match it and remove it",
			Input:       []string{testJwt},
			ExpectSlice: []string{},
			ExpectMatch: testJwt,
			ExpectErr:   false,
		},
		{
			Name:        "When a single JWT and flags are present match it and remove it leaving the args",
			Input:       []string{testJwt, "-e", "--help"},
			ExpectSlice: []string{"-e", "--help"},
			ExpectMatch: testJwt,
			ExpectErr:   false,
		},
		{
			Name:      "When only flags are present expect an error",
			Input:     []string{"-e", "--help"},
			ExpectErr: true,
			ErrVal:    "No valid JWT found",
		},
		{
			Name:      "When empty slice is provided expect an error",
			Input:     []string{},
			ExpectErr: true,
			ErrVal:    "No valid JWT found",
		},
		{
			Name:        "When an improperly formatted JWT is present match it anyway and remove it",
			Input:       []string{testBadJwt},
			ExpectSlice: []string{},
			ExpectMatch: testBadJwt,
			ExpectErr:   false,
		},
		{
			Name:        "When an improperly formatted JWT and flags are present match it anyway and remove it",
			Input:       []string{testBadJwt, "-e", "--help"},
			ExpectSlice: []string{"-e", "--help"},
			ExpectMatch: testBadJwt,
			ExpectErr:   false,
		},
		{
			Name:        "When multiple JWT are present match the last one and remove them all",
			Input:       []string{testJwt, testBadJwt, testJwt},
			ExpectSlice: []string{},
			ExpectMatch: testJwt,
			ExpectErr:   false,
		},
		{
			Name:        "When multiple JWT and flags are present match the last one and remove them all leaving the args",
			Input:       []string{testJwt, "-e", "--help", testBadJwt, testJwt},
			ExpectSlice: []string{"-e", "--help"},
			ExpectMatch: testJwt,
			ExpectErr:   false,
		},
	}

	for _, test := range testcases {

		t.Run(test.Name, func(t *testing.T) {
			gotSlice, gotMatch, gotErr := jwt.FindAndRemoveJwt(test.Input)
			if test.ExpectErr {

				if gotErr == nil {
					t.Logf("received no error but expected error: %s", test.ErrVal)
					t.FailNow()
				}

				if gotErr.Error() != test.ErrVal {
					t.Logf("expected error: %s; got error %s", test.ErrVal, gotErr.Error())
					t.FailNow()
				}

			} else { // not expecting an error

				if gotErr != nil {
					//we got an error we didn't like
					t.Logf("unexpected error encountered: %s", gotErr.Error())
					t.FailNow()
				}

				if test.ExpectMatch != gotMatch {
					t.Logf("invalid match expected string: %s; got string %s", test.ExpectMatch, gotMatch)
					t.Fail()
				}

				assertSlicesMatch(t, test.ExpectSlice, gotSlice)

			}
		})

	}

}

func assertSlicesMatch(t *testing.T, expected []string, actual []string) {

	if len(expected) != len(actual) {

		t.Logf("mismatched slice lengths; expected: %v; got: %v", expected, actual)
		t.FailNow()

	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Logf("mismatch at index %v; expected: %v; got: %v", i, expected[i], actual[i])
			t.Fail()
		}
	}

}
