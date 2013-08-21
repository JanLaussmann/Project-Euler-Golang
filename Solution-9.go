/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a² + b² = c²
For example, 3² + 4² = 9 + 16 = 25 = 5².

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import "fmt"

type triplet struct {
	a, b, c int
}

func main() {
	generator := tripletGenerator()
	for {
		triplet := <-generator
		if triplet.isSearchedFor() {
			fmt.Println(triplet.a * triplet.b * triplet.c)
			break
		}
	}
}

func tripletGenerator() chan triplet {
	channel := make(chan triplet)

	go func() {
		for c := 3; ; c++ {
			for b := 2; b < c; b++ {
				for a := 1; a < b; a++ {
					channel <- triplet{a, b, c}
				}
			}
		}
	}()

	return channel
}

func (t triplet) isSearchedFor() bool {
	return t.a+t.b+t.c == 1000 && t.a*t.a+t.b*t.b == t.c*t.c
}
