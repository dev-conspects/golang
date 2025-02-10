package main

import "fmt"

func main() {
	generator := suqareIntGenerator(7)

	for i := 0; i < 5; i++ {
		fmt.Println(generator())
	}
}

func suqareIntGenerator(startWith int) func() int {
	return func() int {
		res := startWith
		startWith = startWith * startWith
		return res
	}
}
