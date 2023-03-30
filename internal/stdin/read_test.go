package stdin

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
)

type ReadFromTestInput struct {
	Name      string
	PipeFunc  func() (*os.File, *os.File, error)
	IsPipe    bool
	InputVal  string
	ExpectVal string
}

func TestReadFrom(t *testing.T) {
	tests := []ReadFromTestInput{
		{
			Name:      "Happy #1",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  "testing",
			ExpectVal: "testing",
		},
		{
			Name:      "Sample JWT",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			ExpectVal: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		},
		{
			Name:      "Ensure newline is trimmed #1",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  fmt.Sprintln(""),
			ExpectVal: "",
		},
		{
			Name:      "Ensure newline is trimmed #2",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  fmt.Sprintln("test"),
			ExpectVal: "test",
		},
		{
			Name:      "Ensure trailing space is trimmed",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  "testing ",
			ExpectVal: "testing",
		},
		{
			Name:      "Ensure leading space is trimmed",
			PipeFunc:  os.Pipe,
			IsPipe:    true,
			InputVal:  " testing",
			ExpectVal: "testing",
		},
		{
			Name:      "Not Pipe #1", // reads from ./test.txt
			IsPipe:    false,
			ExpectVal: "",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var f *os.File
			if test.IsPipe { // setup a pipe to test with as a mock for stdin
				r, w, err := test.PipeFunc()
				if err != nil {
					t.Log(err)
					t.FailNow()
				}
				_, err = w.WriteString(test.InputVal)
				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				err = w.Close()
				defer func() { r.Close() }()
				if err != nil {
					t.Log(err)
					t.FailNow()
				}
				f = r
			} else { // just read from a file, stdin.readFrom is expected to ignore this
				fileHandle, err := os.OpenFile("test.txt", os.O_RDONLY, fs.FileMode(os.O_RDONLY))
				if err != nil {
					t.Log(err)
					t.FailNow()
				}
				defer func() { fileHandle.Close() }()
				f = fileHandle
			}

			actual := readFrom(f)
			if actual != test.ExpectVal {
				t.Log(fmt.Sprintf("expected: \"%s\"; got: \"%s\"", test.ExpectVal, actual))
				t.FailNow()
			}
		})
	}
}

func TestRead(t *testing.T) { // simple test to make sure read is actually using os.Stdin

	r, w, e := os.Pipe()
	if e != nil {
		t.Logf("unable to setup test pipe, received error: %s", e.Error())
		t.FailNow()
	}

	testVal := "beep boop"
	_, e = w.WriteString(testVal)
	if e != nil {
		t.Logf("unable to write on test pipe, received error: %s", e.Error())
		t.FailNow()
	}

	e = w.Close()
	if e != nil {
		t.Logf("unable to close test pipe, received error: %s", e.Error())
		t.FailNow()
	}

	// ensure Stdin isn't permanently mangled
	save := os.Stdin
	defer func() { os.Stdin = save }()
	os.Stdin = r // temporarily mangle stdin with our own pipe for testing

	actual := Read()

	if actual != testVal {
		t.Logf("read value doesn't appear to be from os.Stdin!")
		t.FailNow()
	}

}
