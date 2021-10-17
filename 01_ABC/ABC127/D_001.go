package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct{ b, c int }

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	bc := make([]pair, m)
	for i := range bc {
		fmt.Fscan(in, &bc[i].c, &bc[i].b)
	}
	sort.Slice(bc, func(i, j int) bool {
		return bc[i].b > bc[j].b
	})

	cnt := 0
	d := make([]int, 0)
	for len(d) < n {
		for i := 0; i < bc[cnt].c; i++ {
			d = append(d, bc[cnt].b)
		}
		cnt++
		if cnt == m {
			break
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d)))

	sum := 0
	if n > len(d) {
		diff := n - len(d)
		for i := 0; i < diff; i++ {
			d = append(d, 0)
		}
	}
	for i := 0; i < n; i++ {
		sum += max(a[i], d[i])
	}
	fmt.Println(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
