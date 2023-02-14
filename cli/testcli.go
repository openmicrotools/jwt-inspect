package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// "" is the default, where no command (such as decode) is included
	// if this line was 'decodeCommand := flag.NewFlagSet("decode", flag.PanicOnError)', you would need to run
	// 'testcli decode' to access the functionality and flags for that command
	decodeCommand := flag.NewFlagSet("", flag.PanicOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --epoch
	/*
		decodeExamplePtr := decodeCommand.String("example", "default result if flag not set", "explanation of "
		// use format that matches decodeCommand.<Type>() for 2nd parameter, so no quotes for bool/int...
		is interpreted:
		testcli <jwt> --example=<setting for example flag> --<flag 2>=<setting for flag2
		For safest usage, use = . `jwt-inspect <jwt> --epoch false` is not equal to `jwt-inspect <jwt> --epoch=false`
	*/
	decodeEpochPtr := decodeCommand.Bool("epoch", false, "Prints the time stamps of the token using epoch format. (Optional)")
	// BoolVar takes in a pointer to store the value to
	// we give it the previous prt for the long flag so as to only need to check one location
	// if multiple flags are passed, the last passed flag will be the value referenced by this var.
	decodeCommand.BoolVar(decodeEpochPtr, "e", false, "Prints the time stamps of the token using epoch format. (Optional)")
	if len(os.Args) < 1 {
		fmt.Println("Usage is 'testcli <jwt> <flags>")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	// command help provided
	case "help":
		fmt.Println("Printing the help documentation!")
	// no command provided option. This is jwt-inspect <jwt>
	default:
		err := decodeCommand.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}

	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	// this still enters with just `testcli jwt` and prints with --epoch=false
	if decodeCommand.Parsed() {
		//Choice flag

		if *decodeEpochPtr {
			fmt.Println("we should provide epoch format")
		} else {
			fmt.Println("we should not provide epoch format")
		}

		// printBool := (*decodeEpochPtr || *decodeEpochPtrShort)
		// Print
		fmt.Printf("decodeEpochPtr: %s,  Print JWT decoded\n",
			fmt.Sprint(*decodeEpochPtr),
		)
	}
}
