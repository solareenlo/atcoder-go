package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n int
	fmt.Scan(&n)

	type pair struct{ x, y int }
	xy := make([][]pair, 20)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		for j := 0; j < a; j++ {
			var x, y int
			fmt.Scan(&x, &y)
			x--
			xy[i] = append(xy[i], pair{x, y})
		}
	}

	maxi := 0
	for bit := 0; bit < 1<<n; bit++ {
		ok := true
		for i := 0; i < n && ok; i++ {
			if bit>>i&1 == 0 {
				continue
			}
			for _, p := range xy[i] {
				if bit>>p.x&1 != p.y {
					ok = false
				}
			}
		}
		if ok {
			maxi = max(maxi, bits.OnesCount(uint(bit)))
		}
	}

	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
