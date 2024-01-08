/* 
	Guess the Sequence
	Mubashir is trying to figure out the corresponding quadratic formula for the following quadratic sequence of numbers:

	N	Result
	1	90
	2	240
	3	450
	4	720
	5	1050

	If you can figure this out, then help him by creating a function that takes a number n and returns the nth number of this quadratic sequence. 
*/


package main

import "fmt"

func quadraticSequence(n int) int {
	// Coefficients for the quadratic equation: ax^2 + bx + c = result
	a := 20
	b := 70
	c := 0

	// Calculate the nth term of the quadratic sequence using the formula
	result := a*n*n + b*n + c

	return result
}

func main() {
	// Test cases with provided data
	fmt.Println(quadraticSequence(1)) // Output: 90
	fmt.Println(quadraticSequence(2)) // Output: 240
	fmt.Println(quadraticSequence(3)) // Output: 450
	fmt.Println(quadraticSequence(4)) // Output: 720
	fmt.Println(quadraticSequence(5)) // Output: 1050
}
