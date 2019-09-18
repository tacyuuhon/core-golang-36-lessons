package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

//
// Run:
// go run main.go -name="Robert"
//
// Help:
// go run main.go --help
func main() {
	fmt.Fprintf(os.Stderr, "os.Args[0]: %s\n", os.Args[0])
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("ip: %s", name)
}
