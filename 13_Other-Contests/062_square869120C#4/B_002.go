package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	res := int(1e16)
	for bit := 0; bit < 1<<n; bit++ {
		var cost, cnt, prev int
		for i := 0; i < n; i++ {
			if bit&(1<<i) == 0 {
				cost += int(math.Max(0, float64(prev+1-a[i])))
				cnt++
				prev++
			}
			prev = int(math.Max(float64(prev), float64(a[i])))
		}
		if bitcnt(bit) == 0 {
			res = int(math.Min(float64(res), float64(cost)))
		}
	}
	fmt.Println(res)
}

func bitcnt(x int) (res int) {
	for x > 0 {
		res += x & 1
		x >>= 1
	}
	return res
}
