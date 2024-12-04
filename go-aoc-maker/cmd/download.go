package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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

	if err := DownloadInputForDay(dayNum); err != nil {
		return err
	}

	if err := DownloadTestInputForDay(dayNum); err != nil {
		return err
	}

	return nil
}

func DownloadInputForDay(day uint64) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day),
		nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create new request for day %d with error %s\n", day, err)
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

func DownloadTestInputForDay(day uint64) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/2024/day/%d", day),
		nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create  new request for day %d with error %s\n", day, err)
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

	s := string(content)

	example := regexp.MustCompile("For example").FindStringIndex(s)
	if example == nil {
		fmt.Fprintf(os.Stderr, "Couldn't find `For example` in the response\n")
		return errors.New("couldn't find `For example` in the response")
	}

	start := regexp.MustCompile("<code>").FindStringIndex(s[example[1]:])
	if start == nil {
		fmt.Fprintf(os.Stderr, "Couldn't find `<code>` in the response\n")
		return errors.New("couldn't find `<code>` in the response")
	}

	startIndex := start[1] + example[1]

	end := regexp.MustCompile("</code>").FindStringIndex(s[startIndex:])
	if end == nil {
		fmt.Fprintf(os.Stderr, "Couldn't find `</code>` in the response\n")
		return errors.New("couldn't find `</code>` in the response")
	}
	
	endIndex := end[0] + startIndex -1

	if err := os.WriteFile("test_data.input", []byte(s[startIndex:endIndex]), os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write test input into file with error: %s\n", err)
		return err
	}

	fmt.Println("The test input was downloaded and saved into test_data.input file")
	return nil
}