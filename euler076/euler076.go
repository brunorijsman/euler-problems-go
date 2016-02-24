package main

import "fmt"

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func countSums(n int) int {
	return countSumsRecurse(n, n)
}

// How many ways can the number N be written as a sum of terms
// a. Where the terms are non-increasing: term(i) <= term(i+1)
// b. Where each term >= 1
// c. Where each term <= maxTerm
//
func countSumsRecurse(n, maxTerm int) int {

	count := 0

	// We can write N as a sum of exactly one term, namely N itself
	//
	if n <= maxTerm {
		count++
	}

	for i := 1; i < n; i++ {

		// We can write N as the following sum:
		//
		// (N-i) xxxxx
		//
		// Where xxxxx is a sum of terms:
		// a. where the terms are non-increasing, hence <= (N-i)
		// b. where each term >= 1
		// c. where each term <= maxTerm
		//
		thisTerm := n - i
		if thisTerm <= maxTerm {
			remainingN := i
			newMaxTerm := min(maxTerm, thisTerm)
			count += countSumsRecurse(remainingN, newMaxTerm)
		}
	}

	return count
}

func main() {
	fmt.Println(countSums(100) - 1)
}
