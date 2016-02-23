package main

import "fmt"

func main() {
	max := 1000
	sum := 0
	for i := 1; i < max; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	fmt.Println(sum)
}
