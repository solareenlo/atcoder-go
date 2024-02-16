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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] = 100 - a[i]
	}
	s := make([]int, n)
	copy(s, a)
	sort.Ints(s)
	for i := 0; i < n; i++ {
		res := lowerBound(s, a[i])
		res++
		fmt.Fprintln(out, res)
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
