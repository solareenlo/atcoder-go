package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	c := make([]int, n)
	s := make([]int, n+1)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	sort.Ints(c)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + c[i]
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var x int
		fmt.Fscan(in, &x)
		fmt.Println(upperBound(s, x) - 1)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
