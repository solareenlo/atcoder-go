package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 200010

var h, f [maxn]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c int
	fmt.Fscan(in, &n, &c)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
	}
	var Q [maxn]int
	Q[0] = 1
	var l, r int
	for i := 2; i <= n; i++ {
		for l < r && k(Q[l+1], Q[l]) < float64(2*h[i]) {
			l++
		}
		f[i] = f[Q[l]] + h[Q[l]]*h[Q[l]] - 2*h[i]*h[Q[l]] + h[i]*h[i] + c
		for l < r && k(i, Q[r]) < k(Q[r], Q[r-1]) {
			r--
		}
		r++
		Q[r] = i
	}
	fmt.Println(f[n])
}

func k(a, b int) float64 {
	return float64(f[a]+h[a]*h[a]-f[b]-h[b]*h[b]) / float64(h[a]-h[b])
}
