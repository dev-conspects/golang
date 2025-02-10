package main

import "fmt"

func main() {

	ch := squareIntGenerator(7, 12)
	for val := range ch {
		fmt.Println(val)
	}
}

func squareIntGenerator(start, end int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		nexVal := start
		for i := start; i <= end; i++ {
			out <- nexVal
			nexVal = nexVal * nexVal
		}
	}()

	return out
}
