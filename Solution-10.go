// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

package main

import "fmt"

func isPrime(in int64) bool {
	var i int64 = 2
	for ; i*i <= in; i++ {
		if in%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var i, sum, lim int64 = 2, 0, 2000000
	for ; i < lim; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
