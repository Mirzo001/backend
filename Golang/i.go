package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(i int) {
	if i%3 == 0 && i%5 == 0 {
		fmt.Printf("\nFizzBuzz")
	} else if i%3 == 0 {
		fmt.Printf("\nFizz")
	} else if i%5 == 0 {
		fmt.Printf("\nBuzz")
	} else {
		fmt.Printf("\n%v", i)
	}
}
