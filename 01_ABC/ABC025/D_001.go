package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x [5][5]int = [5][5]int{}
	for i := range x {
		for j := range x[i] {
			fmt.Scan(&x[i][j])
			x[i][j]--
		}
	}

	dp := make([]int, 1<<25)
	dp[0] = 1
	for bit := 0; bit < 1<<25; bit++ {
		if dp[bit] == 0 {
			continue
		}
		next := bits.OnesCount(uint(bit))
		pos := -1
		for i := range x {
			for j := range x[i] {
				if x[i][j] == next {
					pos = i*5 + j
				}
			}
		}
		for i := range x {
			for j := range x[i] {
				k := i*5 + j
				if (bit>>k)&1 == 1 {
					continue
				}
				if pos != -1 && k != pos {
					continue
				}
				ok := true
				if 0 < i && i < 4 {
					if ((bit>>(k-5))&1)^(bit>>(k+5)&1) == 1 {
						ok = false
					}
				}
				if 0 < j && j < 4 {
					if ((bit>>(k-1))&1)^(bit>>(k+1)&1) == 1 {
						ok = false
					}
				}
				if ok {
					dp[bit|(1<<k)] = (dp[bit|(1<<k)] + dp[bit]) % int(1e9+7)
				}
			}
		}
	}
	fmt.Println(dp[(1<<25)-1])
}
