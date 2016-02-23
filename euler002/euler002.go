package main

import "fmt"

func main() {
	sum := 0
	fib := 1
	prevFib := 1
	for fib <= 4000000 {
		prevFib, fib = fib, prevFib+fib
		if fib%2 == 0 {
			sum += fib
		}
	}
	fmt.Println(sum)
}
