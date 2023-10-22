package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n, m, k int
var a, b []int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &k)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	l, r := 0, int(1e9)
	for l < r {
		md := (l + r) / 2
		if check(md) {
			r = md
		} else {
			l = md + 1
		}
	}
	fmt.Println(l)
}

func check(x int) bool {
	i, j, cnt := 0, 0, 0
	for i < n && j < m {
		if abs(a[i]-b[j]) <= x {
			i++
			j++
			cnt++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return cnt >= k
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
