package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005

var idx int
var top [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	under := make([]int, N)
	p := make([]int, N)
	ans := make([]int, N)
	for i, _ := range ans {
		ans[i] = -1
	}
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		var t int
		if x > top[idx] {
			idx++
			top[idx] = x
			p[x] = 1
			t = idx
		} else {
			t = upper(x)
			p[x] = p[top[t]] + 1
			under[x] = top[t]
			top[t] = x
		}
		if p[x] == k {
			for x > 0 {
				ans[x] = i
				x = under[x]
			}
			for j := t; j < idx; j++ {
				top[j] = top[j+1]
			}
			idx--
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func upper(t int) int {
	l, r, res := 1, idx, 1
	for l < r {
		mid := (l + r) >> 1
		if top[mid] > t {
			r = mid
		} else {
			l, res = mid+1, mid+1
		}
	}
	return res
}
