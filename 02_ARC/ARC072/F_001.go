package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	l := 1
	s := 0
	w := make([]int, 500500)
	v := make([]float64, 500500)
	c := 0.0
	r := 0
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		for y+s > m {
			p := min(y+s-m, w[l])
			s -= p
			w[l] -= p
			c -= v[l] * float64(p)
			if w[l] == 0 {
				l++
			}
		}
		r++
		w[r] = y
		v[r] = float64(x)
		s += y
		c += float64(x * y)
		for l < r && v[r] < v[r-1] {
			r--
			v[r] = (v[r]*float64(w[r]) + v[r+1]*float64(w[r+1])) / float64(w[r]+w[r+1])
			w[r] += w[r+1]
		}
		fmt.Fprintln(out, c/float64(m))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
