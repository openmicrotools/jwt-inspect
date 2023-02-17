package cmdinput

import (
	"flag"
	"fmt"
	"os"

	"github.com/openmicrotools/jwt-inspect/pkg/jwt"
	"github.com/openmicrotools/jwt-inspect/pkg/stdin"
)

const HelpMsg = `
jwt-inspect [flags] token
Requires a JWT (JSON Web Token) either as an argument to the cli or piped in on stdin. If a JWT cannot be identified in either location it will error.

`

func exitUsage(exitVal int, fs *flag.FlagSet, messages ...string) {
	fmt.Print(HelpMsg)
	if exitVal == 0 { // print to stdout because this isn't an error condition
		fs.SetOutput(os.Stdout)
	}
	fs.Usage()
	fmt.Println("") // insert empty line
	for _, m := range messages {
		fmt.Println(m)
	}
	os.Exit(exitVal)
}

func registerBool(fs *flag.FlagSet, name string, defaultVal bool, usage string) *bool {

	var flagVal *bool

	flagVal = fs.Bool(name, defaultVal, usage)
	fs.BoolVar(flagVal, name[:1], defaultVal, usage)

	return flagVal

}

func HandleInput(args []string) (string, bool) {
	jwtInput := stdin.Read()

	defaultCmd := flag.NewFlagSet("", flag.ExitOnError)
	printEpoch := registerBool(defaultCmd, "epoch", false, "Output any numeric date fields in the token using epoch format instead of as a formatted human readable string. (Optional)")
	printHelp := registerBool(defaultCmd, "help", false, "Print the usage information for jwt-inspect. (Optional)")

	if jwtInput == "" {
		var err error
		args, jwtInput, err = jwt.FindAndRemoveJwt(args)
		if err != nil {
			exitUsage(1, defaultCmd, err.Error())
		}
	}

	err := defaultCmd.Parse(args)
	if err != nil {
		exitUsage(2, defaultCmd, fmt.Sprintf("unable to parse arguments due to error: %s", err.Error()))
	}

	if !defaultCmd.Parsed() || printEpoch == nil || printHelp == nil {
		exitUsage(3, defaultCmd, "command line args not properly parsed for unknown reason")
	}

	if *printHelp {
		exitUsage(0, defaultCmd)
	}
	return jwtInput, *printEpoch
}
