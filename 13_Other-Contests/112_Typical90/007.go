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

	var N int
	fmt.Fscan(in, &N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var b int
		fmt.Fscan(in, &b)
		i := lowerBound(a, b)
		ans := int(2e9)
		if i < N {
			ans = a[i] - b
		}
		i--
		if i >= 0 {
			ans = min(ans, b-a[i])
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
