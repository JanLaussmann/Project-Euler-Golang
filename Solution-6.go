// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func main() {
	var a, b int64 = 1, 100
	fmt.Println(squareOfSums(a, b) - sumOfSquares(a, b))
}

func sumOfSquares(a, b int64) int64 {
	var x int64 = 0
	for ; a <= b; a++ {
		x += a * a
	}
	return x
}

func squareOfSums(a, b int64) int64 {
	var x int64 = 0
	for ; a <= b; a++ {
		x += a
	}
	return x * x
}