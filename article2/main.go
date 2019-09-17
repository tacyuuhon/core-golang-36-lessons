package main

import "fmt"

var name string = "Go"

func init() {
	fmt.Println("init func ....")
}

func main() {
	fmt.Printf("Hello, %s!\n", name)
}
