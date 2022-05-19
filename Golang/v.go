package main

import "fmt"

func main() {
	n := 4
	// Read n from input
	DisplayMinimumNumberFunction(n)
}

// https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number
func DisplayMinimumNumberFunction(n int) {
	var i int

	for i = 1; i <= n-2; i++ {
		fmt.Printf("min(int, ")
	}
	fmt.Printf("min(int, int)")

	for i = 2; i <= n-1; i++ {
		fmt.Printf(")")
	}
}
