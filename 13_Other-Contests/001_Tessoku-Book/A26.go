package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var q int
	fmt.Fscan(in, &q)

	for q > 0 {
		q--
		var x int
		fmt.Fscan(in, &x)
		if isPrime(x) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
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
