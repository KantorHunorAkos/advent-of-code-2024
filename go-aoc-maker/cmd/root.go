package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "aocmaker",
	Short: "Makes an advent of code day template",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occured while executing aocmaker '%s'\n", err)
		os.Exit(1)
	}
}