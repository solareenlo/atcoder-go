package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	A := make([]int, 1<<17)
	for ; m > 0; m-- {
		var a int
		fmt.Fscan(in, &a)
		idx := upperBound(A[:n], -a)
		A[idx] = -a
		if idx < n {
			fmt.Println(idx + 1)
		} else {
			fmt.Println(-1)
		}
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
