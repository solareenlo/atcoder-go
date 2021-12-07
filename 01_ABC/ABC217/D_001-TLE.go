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

	var l, q int
	fmt.Fscan(in, &l, &q)

	m := map[int]bool{}
	m[0] = true
	m[l] = true

	for i := 0; i < q; i++ {
		var c, x int
		fmt.Fscan(in, &c, &x)
		if c == 1 {
			m[x] = true
		} else {
			keys := make([]int, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			pos := lowerBound(keys, x)
			r := keys[pos]
			l := keys[pos-1]
			fmt.Fprintln(out, r-l)
		}
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
