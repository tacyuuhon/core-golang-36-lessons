package main

import (
	"flag"
	"fmt"
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
	flag.Parse()
	fmt.Printf("ip: %s", name)
}
