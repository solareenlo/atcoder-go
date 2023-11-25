package main

import (
	"fmt"
	"math"
)

func main() {
	var output func(int, int, int)
	output = func(a, b, c int) {
		M := a + b - c
		fmt.Printf("3 %d\n", M)
		for c > 0 {
			c--
			a--
			b--
			fmt.Println("1 2")
		}
		for a > 0 {
			a--
			fmt.Println("1 3")
		}
		for b > 0 {
			b--
			fmt.Println("2 3")
		}
	}

	var K int
	fmt.Scan(&K)
	for a := int(math.Ceil(math.Sqrt(float64(K)))); ; a++ {
		for b := (K + a - 1) / a; b <= a; b++ {
			c := int(math.Sqrt(float64(a*b - K)))
			if c > b {
				continue
			}
			if a*b-c*c == K {
				output(a, b, c)
				return
			}
		}
	}
}
