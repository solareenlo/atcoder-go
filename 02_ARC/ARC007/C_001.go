package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	c := [12]int{}
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			c[0] |= 1 << i
		}
	}

	for i := 0; i < n-1; i++ {
		c[i+1] = c[i]*2%(1<<n) + (c[i] >> (n - 1))
	}

	ans := 12
	for bit := 0; bit < 1<<n; bit++ {
		ok := 0
		cnt := bits.OnesCount(uint(bit))
		for i := 0; i < n; i++ {
			if (bit>>i)&1 != 0 {
				ok |= c[i]
			}
		}
		if ok+1 == 1<<n {
			ans = min(ans, cnt)
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
