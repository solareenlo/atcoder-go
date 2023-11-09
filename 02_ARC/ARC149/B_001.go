package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300005

	var a, b, c [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[b[i]])
		b[i] = 1
	}
	c[0] = a[1]
	s := 1
	for i := 2; i <= n; i++ {
		if a[i] > c[s-1] {
			c[s] = a[i]
			s++
		} else {
			c[lowerBound(c[:s], a[i])] = a[i]
		}
	}
	fmt.Println(s + n)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
