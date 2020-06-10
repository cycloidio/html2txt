package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"jaytaylor.com/html2text"
)

func main() {
	os.Exit(run())
}

func run() int {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error when reading stdin: %v\n", err)
		return 1
	}

	txt, err := html2text.FromString(string(input), html2text.Options{PrettyTables: false})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error when generating txt from html: %v\n", err)
		return 1
	}

	fmt.Fprintln(os.Stdout, txt)
	return 0
}
