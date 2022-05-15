package main

import "fmt"

func main() {
	var a, b int = 3, 4

	// fmt.Scanf("%d", &a)
	// fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println(a, "is odd: ", isOdd(a))
	fmt.Println(b, "is even: ", isEven(b))
}

func isEven(num int) bool {
	return num%2 == 0
}

func isOdd(num int) bool {
	return num%2 == 1
}
