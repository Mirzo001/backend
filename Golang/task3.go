package main

import (
	"fmt"
	"math"
)

func main() {
	var r float32 = 10.04

	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {
	area = 4*r*r - math.Pi*r*r
	return area

}
