package cmdinput

import (
	"flag"
	"fmt"
	"os"

	"github.com/openmicrotools/jwt-inspect/internal/stdin"
)

// define some error code constants
const (
	_ = iota // we don't want any 0 error codes
	errJwtNotFound
	errParseArgs
)

// define the basic help message as a constant so it's easy to update everywhere
const helpMsg = `
jwt-inspect [flags] token
Requires a JWT (JSON Web Token) either as an argument to the cli or piped in on stdin. If a JWT cannot be identified in either location it will error.

`

// HandleInput looks for a valid JWT in stdin or from the cli args, handles flags and returns the input and relevant options.
// Depending on input it may also error out and exit or print a usage message and exit
func HandleInput() (string, bool) {
	return handleInput(os.Args) // call to more testable private function
}

func handleInput(args []string) (string, bool) {
	args = args[1:] //args[0] is the name of the command being run

	jwtInput := stdin.Read() // attempt to read a JWT from stdin

	// Setting up our basic flags
	defaultCmd := flag.NewFlagSet("", flag.ExitOnError)
	printEpoch := registerBool(defaultCmd, "epoch", false, "Output any numeric date fields in the token using epoch format instead of as a formatted human readable string. (Optional)")
	printHelp := registerBool(defaultCmd, "help", false, "Print the usage information for jwt-inspect. (Optional)")

	var inputErr error  // capture and store the inputErr value for later use within the outer scope
	if jwtInput == "" { // if we didn't get anything on stdin let's look for an argument with a JWT
		args, jwtInput, inputErr = findAndRemoveJwt(args) // try to peel off a JWT from our list of args to make flag parsing more smooth; without this flag parsing will get tripped up on the first non-flag argument and quit instead of parsing flags which may come after the jwt
	}

	err := defaultCmd.Parse(args) // parse cli flags
	if err != nil {
		exitUsage(errParseArgs, defaultCmd, fmt.Sprintf("unable to parse arguments due to error: %s", err.Error()))
	}

	if !defaultCmd.Parsed() || printEpoch == nil || printHelp == nil { // check that parse was OK and that no pointers are still nil
		exitUsage(errParseArgs, defaultCmd, "command line args not properly parsed for unknown reason")
	}

	if *printHelp { // if we're being asked to print help lets exit with usage but not with an error
		exitUsage(0, defaultCmd)
	}

	if inputErr != nil { // we check for input errors after parsing and running through possible help scenario to avoid errors when only the -h flag is provided
		exitUsage(errJwtNotFound, defaultCmd, inputErr.Error())
	}

	return jwtInput, *printEpoch // looks successful so far, just return our jwt and options
}

func exitUsage(exitVal int, fs *flag.FlagSet, messages ...string) {
	fmt.Print(helpMsg) // print our basic usage information
	if exitVal == 0 {  // print to stdout because this isn't an error condition
		fs.SetOutput(os.Stdout)
	}
	fs.Usage()                   // print usage for each flag
	fmt.Println("")              // insert empty line
	for _, m := range messages { // if there is any other input range over it and print it all
		fmt.Println(m)
	}
	os.Exit(exitVal) // terminate execution
}

func registerBool(fs *flag.FlagSet, name string, defaultVal bool, usage string) *bool {

	flagVal := fs.Bool(name, defaultVal, usage)      // register the long flag
	fs.BoolVar(flagVal, name[:1], defaultVal, usage) // register the short flag as well

	return flagVal

}
