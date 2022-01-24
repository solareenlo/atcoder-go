package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	prime := false
	if n == 1 {
		prime = false
	} else if isPrime(n) {
		prime = true
	} else {
		s := strconv.Itoa(n)
		sum := 0
		for i := 0; i < len(s); i++ {
			sum += int(s[i] - '0')
		}
		if n%2 != 0 {
			if n%5 != 0 {
				if sum%3 != 0 {
					prime = true
				}
			}
		}
	}

	if prime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
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
