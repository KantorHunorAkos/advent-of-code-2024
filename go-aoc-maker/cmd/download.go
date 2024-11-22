package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use: "download",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Downloads the input for day %s\n", args[0])
		DownloadForDay(args[0])
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func DownloadForDay(day string) error {
	dayNum, err := strconv.ParseUint(day, 10, strconv.IntSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse day as number %s with error %s\n", day, err)
		return err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", dayNum),
		nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create  new request for day %s with error %s\n", day, err)
		return err
	}
	
	aocSession := os.Getenv("AOC_SESSION")
	req.AddCookie(&http.Cookie{Name: "session", Value: aocSession})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading the body failed: %s\n", err)
		return err
	}

	pleaseResponse := []byte("Puzzle inputs differ by user.  Please log in to get your puzzle input.")
	if bytes.Contains(content, pleaseResponse) {
		fmt.Fprint(os.Stderr, "Please log in to get your puzzle input.\n")
		return err
	}

	if err := os.WriteFile("data.input", content, os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write input into file with error: %s\n", err)
		return err
	}

	fmt.Println("The input was downloaded and saved into data.input file")
	return nil
}