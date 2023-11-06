package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var s, q, x [3030]int
var f [3030][3030]int

func S(p, q int) int { return s[p+q-1] - s[q-1] }
func V(p, q int) int { return S(p, q) + f[p][q] }

func cal(k, i, Z int) {
	for j, h, t := 1, 1, 0; j <= n+1-k; j++ {
		for h <= t && V(k, q[t]) >= V(k, j) {
			t--
		}
		t++
		q[t] = j
		for h < t && q[h] < j+k-i {
			h++
		}
		if j+k > i {
			f[i][j+k-i] = max(f[i][j+k-i], S(i, j+k-i)-Z-V(k, q[h]))
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B, C, D int
	fmt.Fscan(in, &n, &A, &B, &C, &D)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
		s[i] = s[i-1] + x[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n+1-i; j++ {
			f[i][j] = max(x[j]-f[i-1][j+1], x[i+j-1]-f[i-1][j])
		}
		cal(i-min(i, B), i, A)
		cal(i-min(i, D), i, C)
	}
	fmt.Println(f[n][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
