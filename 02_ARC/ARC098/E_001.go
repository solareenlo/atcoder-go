package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, K, Q int
	fmt.Fscan(in, &n, &K, &Q)

	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := 0
	r := 1000_000_007
	for l < r {
		mid := (l + r) / 2
		fl := false
		for i := 1; i <= n; i++ {
			p := a[i]
			k := 0
			tmp := 0
			s := 0
			for j := 1; j <= n+1; j++ {
				if a[j] < p || j == n+1 {
					if j-k-1 >= K {
						tmp += min(j-k-1-K+1, j-k-1-s)
					}
					k = j
					s = 0
				} else if a[j] > p+mid {
					s++
				}
			}
			if tmp >= Q {
				fl = true
				break
			}
		}
		if fl {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
