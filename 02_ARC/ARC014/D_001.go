package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, n, m int
	fmt.Fscan(in, &a, &n, &m)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
		l[i]--
	}

	s := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		s[i] = l[i+1] - l[i] - 1
	}
	sort.Ints(s)

	r := make([]int, n)
	for i := 0; i < n-1; i++ {
		r[i+1] = r[i] + s[i]
	}

	ans := 0
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		ans = min(x, l[0]) + min(y, a-l[n-1]-1)
		k := lowerBound(s, x+y)
		ans += r[k] + (n-k-1)*(x+y)
		fmt.Fprintln(out, ans+n)
	}
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
