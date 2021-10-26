package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	ok, ng := 1<<60, -1
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			cnt += max(a[i]-mid/b[i], 0)
		}
		if cnt > k {
			ng = mid
		} else {
			ok = mid
		}
	}

	fmt.Println(ok)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
