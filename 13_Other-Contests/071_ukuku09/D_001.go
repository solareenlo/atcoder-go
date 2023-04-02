package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, q int
var s string
var w [100005]int
var h [20][100005]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &q, &s)
	manacher()
	for i := 1; i < 17; i++ {
		w[1<<i]++
	}
	for i := 1; i <= n; i++ {
		w[i] += w[i-1]
	}
	for i := 0; i < 17; i++ {
		for j := 0; j < n; j++ {
			if j+(1<<i) < n {
				h[i+1][j] = max(h[i][j], h[i][j+(1<<i)])
			} else {
				h[i+1][j] = h[i][j]
			}
		}
	}
	for q > 0 {
		q--
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		l, r := 0, n/2+1
		for r-l > 1 {
			m := (l + r) / 2
			if c(a, b, m) {
				l = m
			} else {
				r = m
			}
		}
		fmt.Fprintln(out, l*2+1)
	}
}

func manacher() {
	i, j := 0, 0
	for i < n {
		for i-j >= 0 && i+j < n && s[i-j] == s[i+j] {
			j++
		}
		h[0][i] = j
		k := 1
		for i-k >= 0 && i+k < n && k+h[0][i-k] < j {
			h[0][i+k] = h[0][i-k]
			k++
		}
		i += k
		j -= k
	}
}

func c(l, r, x int) bool {
	l += x
	r -= x
	if l > r {
		return false
	}
	y := w[r-l+1]
	return max(h[y][l], h[y][r-(1<<y)+1])-1 >= x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
