package stdin

import (
	"io"
	"os"
	"strings"
)

// Read will check if there's a named pipe available on stdin and return a string of that data or "" if anything goes wrong
func Read() string {
	return readFrom(os.Stdin)
}

// easier to test private function
// will return a value if from a pipe
// returns empty string from any non-pipe file though
func readFrom(f *os.File) string {
	stat, err := f.Stat() // get file stats on stdin
	if err != nil {
		return ""
	}
	if (stat.Mode() & os.ModeNamedPipe) != 0 { // check that stdin is of type "named pipe" and attempt to read from it if it is a pipe
		bytes, err := io.ReadAll(f) // lets try to read from stdin
		if err == nil || len(bytes) > 0 {
			return strings.TrimSpace(string(bytes)) // convert to string and remove whitespace
		}
	}
	return ""
}
