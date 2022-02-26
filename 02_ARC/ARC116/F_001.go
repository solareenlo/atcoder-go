package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200200

var (
	m  int
	a  = make([]int, N)
	r  = make([]int, N)
	l  = make([]int, N)
	n  = make([]int, N)
	st = make([]int, N)
)

func calc(L, R int) int {
	if L == R {
		return a[L]
	}
	mid := (L + R) >> 1
	if (r[m]-m)&1 != 0 {
		return max(a[mid], min(a[mid-1], a[mid+1]))
	}
	return min(a[mid], max(a[mid-1], a[mid+1]))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &m)

	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &n[i])
		l[i] = r[i-1] + 1
		r[i] = r[i-1] + n[i]
		for j := l[i]; j <= r[i]; j++ {
			fmt.Fscan(in, &a[j])
		}
	}

	res := 0
	t := 0
	for i := 1; i <= m; i++ {
		if n[i]&1 != 0 {
			res += calc(l[i], r[i])
		} else {
			a := calc(l[i], r[i]-1)
			b := calc(l[i]+1, r[i])
			if a > b {
				a, b = b, a
			}
			res += a
			t++
			st[t] = b - a
		}
	}
	tmp := st[1 : t+1]
	sort.Ints(tmp)

	for i := t; i >= 1; i -= 2 {
		res += st[i]
	}
	fmt.Println(res)
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
