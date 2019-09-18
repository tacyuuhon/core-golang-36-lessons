package main

import (
	"flag"

	hello "github.com/tacyuuhon/core-golang-36-lessons/article8/hello"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello.Hello(name)
}
