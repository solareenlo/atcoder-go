package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, q int
	fmt.Fscan(in, &a, &b, &q)
	s := make([]int, a+2)
	s[0], s[a+1] = -(1 << 60), 1<<60
	for i := 1; i <= a; i++ {
		fmt.Fscan(in, &s[i])
	}
	t := make([]int, b+2)
	t[0], t[b+1] = -(1 << 60), 1<<60
	for i := 1; i <= b; i++ {
		fmt.Fscan(in, &t[i])
	}

	var x int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x)

		j := sort.SearchInts(s, x)
		k := sort.SearchInts(t, x)
		sl, sr := abs(x-s[j-1]), abs(x-s[j])
		tl, tr := abs(x-t[k-1]), abs(x-t[k])

		res := 1 << 60
		res = min(res, max(sl, tl), max(sr, tr))
		sd, td := min(sl, sr), min(tl, tr)
		res = min(res, min(sd, td)*2+max(sd, td))

		fmt.Println(res)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
