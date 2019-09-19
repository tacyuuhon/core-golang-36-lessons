package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero-0", 1: "one-1", 2: "two-2"}
	fmt.Printf("The element is %q\n", container[1])
}
