package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}
