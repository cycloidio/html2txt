package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"jaytaylor.com/html2text"
)

func main() {
	bhtml, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("Unexpected error when reading stdin: %v", err)
		return
	}
	txt, err := html2text.FromString(string(bhtml), html2text.Options{PrettyTables: false})
	if err != nil {
		log.Printf("Unexpected error when generating txt from html: %v", err)
		return
	}
	_, err = fmt.Fprintln(os.Stdout, txt)
	if err != nil {
		log.Printf("Unexpected error when writing to stdout: %v", err)
		return
	}
}
