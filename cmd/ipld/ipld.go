package main

import (
	"os"

	"github.com/ipld/go-ipldtool"
)

func main() {
	code, _ := app.Main(os.Args, os.Stdin, os.Stdout, os.Stderr)
	os.Exit(code)
}
