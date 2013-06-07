package main

import "fmt"

func main() {
	a, b, c, sum := 1, 1, 0, 0
	for {
		if c > 4000000 {
			break
		}
		if c % 2 == 0 {
			sum += c
		}
		c = a + b
		a, b = c, a
	}
	fmt.Println(sum)
}