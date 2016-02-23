package main

import "fmt"

func factor(n int) (factors []int) {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return append(factor(n/i), i)
		}
	}
	return []int{n}
}

func main() {
	fmt.Println(factor(600851475143)[0])
}
