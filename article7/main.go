package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {

}

//
// Run:
// go run main.go -name="Robert"
//
// Help:
// go run main.go --help
func main() {
	var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("ip: %s", name)
}
