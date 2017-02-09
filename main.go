package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/adamdecaf/gointerp/repl"
)

func main() {
	r := flag.Bool("repl", false, "Start the repl")
	flag.Parse()

	if r != nil && *r {
		fmt.Println("Starting gointerp repl")
		repl.Start(os.Stdin, os.Stdout)
	}
}
