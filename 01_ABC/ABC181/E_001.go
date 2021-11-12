package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}
	sort.Ints(h)

	l := make([]int, n/2+1)
	r := make([]int, n/2+1)
	for i := 0; i < n-1; i += 2 {
		l[i/2+1] = l[i/2] + (h[i+1] - h[i])
	}
	for i := n - 2; i > 0; i -= 2 {
		r[i/2] = r[i/2+1] + (h[i+1] - h[i])
	}

	res := 1 << 60
	for i := 0; i < m; i++ {
		var w int
		fmt.Fscan(in, &w)
		x := lowerBound(h, w)
		if x&1 != 0 {
			x ^= 1
		}
		res = min(res, l[x/2]+r[x/2]+abs(w-h[x]))
	}

	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
