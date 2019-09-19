package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero-0", 1: "one-1", 2: "two-2"}
	fmt.Printf("The element is %q\n", container[1])
	value, ok := interface{}(12.3).(float64)

	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println(string(-1))
	fmt.Println(string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))

	{
		type MyString1 = string
		var myString1 MyString1 = "abc"
		value, ok := interface{}(myString1).(string)
		fmt.Printf("MyString1 val:%v, is string: %v\n", value, ok)
	}

	{
		type MyString2 string
		var myString2 MyString2 = "abc"
		value, ok := interface{}(myString2).(string)
		fmt.Printf("MyString2 val:%v, is string: %v\n", value, ok)
	}

}
