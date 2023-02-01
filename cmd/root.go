/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jwt-inspect",
	Short: "A CLI to inspect JWT",
	Long: `
jwt-inspect is a CLI tool to inspect JWT. TODO: write longer description`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		jwt := args[0]
		isPrettyPrint, _ := cmd.Flags().GetBool("pretty-print")

		s, _ := decodeJwt(jwt, isPrettyPrint)
		fmt.Println(s)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jwt-inspect.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("pretty-print", "p", true, "pretty print the timestamps defaulting in RFC1123 format")
}

