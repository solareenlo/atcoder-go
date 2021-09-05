package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	sum, res := 1, 0
	for r, l := 0, 0; r < n; r++ {
		sum *= s[r]
		for sum > k && l <= r {
			sum /= s[l]
			l++
		}
		res = int(math.Max(float64(res), float64(r-l+1)))
	}
	if sum != 0 {
		fmt.Println(res)
	} else {
		fmt.Println(n)
	}
}
