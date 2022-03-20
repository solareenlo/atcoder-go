package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqr := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqr; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
