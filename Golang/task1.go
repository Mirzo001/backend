package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Please enter value of a:")
	fmt.Scanf("%v\n", &a)

	fmt.Println("Please enter value of b:")
	fmt.Scanf("%v", &b)

	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)
}
