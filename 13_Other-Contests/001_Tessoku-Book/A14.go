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
	c := make([]int, n)
	d := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	for i := range d {
		fmt.Fscan(in, &d[i])
	}
	p := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p = append(p, a[i]+b[j])
		}
	}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			q = append(q, c[i]+d[j])
		}
	}
	sort.Ints(q)

	for i := 0; i < n*n; i++ {
		pos := lowerBound(q, k-p[i])
		if pos < len(q) && q[pos] == k-p[i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
