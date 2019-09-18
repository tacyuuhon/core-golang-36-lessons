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
//go run main.go -name="Robert"
//
func main() {
	flag.Parse()
	fmt.Printf("ip: %s", name)
}
