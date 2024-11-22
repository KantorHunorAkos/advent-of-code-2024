package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var makerCmd = &cobra.Command{
	Use: "make",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Mekes skeleton and downloads the input for day %s\n", args[0])
		makeFullDay(args[0])
	},
}

func init() {
	rootCmd.AddCommand(makerCmd)
}

func makeFullDay(day string) {
	if err := MakeSkeletonForDay(day); err != nil {
		return
	}

	if err := os.Chdir(fmt.Sprintf("day-%s", day)); err != nil {
		return
	}

	if err := DownloadForDay(day); err != nil {
		return
	}

	goModInitCmd := exec.Command("go", "mod", "init", fmt.Sprintf("day-%s", day))
	stdout, err := goModInitCmd.Output()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while running `go mod init day-%s`: %s", day, err)
		return
	}
	fmt.Printf("Successfully ran command `go mod init day-%s`\n%s", day, stdout)

	if err := os.Chdir(".."); err != nil {
		return
	}
}