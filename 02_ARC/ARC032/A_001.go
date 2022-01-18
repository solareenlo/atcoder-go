package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n+1; i++ {
		sum += i
	}

	if isPrime(sum) {
		fmt.Println("WANWAN")
	} else {
		fmt.Println("BOWWOW")
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
