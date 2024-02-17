package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1005

var v [N]bool
var vt [N][]int
var ans []int

func f(p int) {
	v[p] = true
	for i := 0; i < len(vt[p]); i++ {
		t := vt[p][i]
		if !v[t] {
			f(t)
		}
	}
	ans = append(ans, p)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tp, bt, le, ri [N]int
	var a [105][105]int

	var n, w, h int
	fmt.Fscan(in, &n, &w, &h)
	for i := 0; i < n; i++ {
		tp[i] = h
		bt[i] = -1
		le[i] = w
		ri[i] = -1
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var t int
			fmt.Fscan(in, &t)
			t--
			a[x][y] = t
			if t >= 0 {
				tp[t] = min(tp[t], y)
				bt[t] = max(bt[t], y)
				le[t] = min(le[t], x)
				ri[t] = max(ri[t], x)
			}
		}
	}
	for i := 0; i < n; i++ {
		for x := le[i]; x <= ri[i]; x++ {
			for y := tp[i]; y <= bt[i]; y++ {
				t := a[x][y]
				if t != -1 && t != i {
					vt[t] = append(vt[t], i)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if !v[i] {
			f(i)
		}
	}
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fprintln(out, ans[i]+1)
		} else {
			fmt.Fprintf(out, "%d ", ans[i]+1)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
