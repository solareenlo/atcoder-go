package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	li := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		li[i] = upperBound(a, a[i]+d) - 1
	}
	s[0] = li[0]
	for i := 0; i < n-1; i++ {
		s[i+1] = s[i] + li[i+1]
	}
	c := 0
	for i := 0; i < n; i++ {
		c += s[li[i]] - s[i] - li[i]*(li[i]-i)
	}
	fmt.Println(c)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
