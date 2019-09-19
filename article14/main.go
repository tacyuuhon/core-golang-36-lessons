package main

import "fmt"

func main() {
	{
		aMap := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		}

		k := "two"
		v, ok := aMap[k]
		if ok {

			fmt.Printf("The element of key %q: %d\n", k, v)
		} else {
			fmt.Println("Not fount!")
		}
	}

	{
		open := func() {
			fmt.Println("Open func")
		}

		running := func() {
			fmt.Println("Running func")
		}

		stop := func() {
			fmt.Println("Stop func")
		}

		aMap := map[string]func(){
			"open":    open,
			"running": running,
			"stop":    stop,
		}

		k := "running"
		v, ok := aMap[k]
		if ok {
			v()
		} else {
			fmt.Println("Not fount!")
		}
	}

}
