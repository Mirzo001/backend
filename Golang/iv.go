package main

import (
	"fmt"
)

func main() {
	fmt.Println(MySquareRoot(10, 2))
}

// func MySquareRoot(num, precision uint) (result float64) {
// 	// DO NOT USE math.Sqrt!

// 	//
// 	// WRITE YOUR CODE HERE
// 	//

// 	return
// }

func MySquareRoot(number int, precision uint) float64 {
	// var start int = 0
	// var end int = number
	var mid int

	// variable to store the result
	var result float64

	// for computing integral part
	// of square root of number
	for start, end := int(0), number; start <= end; {
		mid = (start + end) / 2
		if mid*mid == number {
			result = float64(mid)
			break
		}

		// incrementing start if integral
		// part lies on right side of the mid
		if mid*mid < number {
			start = mid + 1
			result = float64(mid)
		} else {
			end = mid - 1
		}
	}

	// For computing the fractional part
	// of square root upto given precision
	var increment float64 = 0.1
	for i := uint(0); i < precision; i++ {
		// while (result * result <= number) {
		// 	result += increment;
		// }
		for {

			if result*result <= float64(number) {
				result += increment

			} else {
				break
			}

		}

		// loop terminates when result * result > number
		result = result - increment
		increment = increment / 10
	}
	return result
}
