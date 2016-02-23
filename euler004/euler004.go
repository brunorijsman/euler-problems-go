package main

import (
	"fmt"
	"strconv"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	m := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			if isPalindrome(i * j) {
				m = max(m, i*j)
			}
		}
	}
	fmt.Println(m)
}
