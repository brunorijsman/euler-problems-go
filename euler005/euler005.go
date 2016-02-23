package main

import "fmt"

const max = 20

func main() {
	for i := max; ; i += max {
		found := true
		for j := 2; j < max; j++ {
			if i%j != 0 {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i)
			break
		}
	}
}
