package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := math.Sqrt(1e9)
	for i := 0; i < int(s)+2; i++ {
		if i*i == n {
			fmt.Println(n)
			return
		} else if i*i > n {
			fmt.Println((i - 1) * (i - 1))
			return
		}
	}
	fmt.Println(0)
}
