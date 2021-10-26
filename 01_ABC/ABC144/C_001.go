package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n - 1)
	} else {
		mini := 1 << 60
		num := 0
		sqr := int(math.Sqrt(float64(n)))
		for i := 2; i < sqr+1; i++ {
			if n%i == 0 {
				if i+n/i < mini {
					mini = min(mini, i+n/i)
					num = i
				}
			}
		}
		fmt.Println(num + n/num - 2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
