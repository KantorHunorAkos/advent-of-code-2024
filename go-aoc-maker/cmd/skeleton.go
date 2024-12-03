package cmd

import (
	"github.com/spf13/cobra"
	_ "embed"
	"fmt"
	"os"
	"strconv"
)

var skeletonCmd = &cobra.Command{
	Use: "skeleton",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating a skeleton for day %s\n", args[0])
		MakeSkeletonForDay(args[0])
	},
}

func init() {
	rootCmd.AddCommand(skeletonCmd)
}

//go:embed templates/main.go
var mainTempl []byte
//go:embed templates/part1.go
var part1Templ []byte
//go:embed templates/part2.go
var part2Templ []byte
//go:embed templates/test.go
var testTempl []byte
func MakeSkeletonForDay(day string) error {
	dayNum, err := strconv.ParseUint(day, 10, strconv.IntSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse day as number %s with error %s\n", day, err)
		return err
	}

	dirName := fmt.Sprintf("day-%d", dayNum)

	fmt.Printf("Creating directory %s\n", dirName)
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory %s with error %s\n", dirName, err)
		return err
	}
	fmt.Printf("Created directory %s\n", dirName)

	fmt.Printf("Changing working directory to %s\n", dirName)
	if err := os.Chdir(dirName); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change working directory to %s with error %s\n", dirName, err)
		return err
	}

	if err := createGoFile("main.go", dirName, mainTempl); err != nil {
		return err
	}
	if err := createGoFile("part1.go", dirName, part1Templ); err != nil {
		return err
	}
	if err := createGoFile("part2.go", dirName, part2Templ); err != nil {
		return err
	}
	if err := createGoFile(fmt.Sprintf("%s_test.go", dirName), dirName, testTempl); err != nil {
		return err
	}
	
	fmt.Printf("Leaving directory %s\n", dirName)
	if err := os.Chdir(".."); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to leave working directory with error %s\n", err)
		return err
	}
	fmt.Printf("Successfully created skeleton for day %d\n", dayNum)
	return nil
}

func createGoFile(fileName string, dirName string, template []byte) error {
	fmt.Printf("Creating %s file in directory %s\n", fileName, dirName)
	if err := os.WriteFile(fileName, template, 0666); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create %s file with error %s\n", fileName, err)

		if err := os.Chdir(".."); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to change working directory to parent with error %s\n", err)
			return err
		}

		if err := os.RemoveAll(dirName); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove directory with error %s\n", err)
			return err
		}

		fmt.Printf("Succsessfully removed directory %s\n", dirName)
		return err
	}
	fmt.Printf("Created %s file in directory %s\n", fileName, dirName)
	return nil
}