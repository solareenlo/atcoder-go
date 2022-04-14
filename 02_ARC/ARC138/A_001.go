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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := k - 1; i >= 1; i-- {
		if a[i] > a[i+1] {
			a[i] = a[i+1]
		}
	}

	ans := 1 << 60
	for i := k + 1; i <= n; i++ {
		p := lowerBound(a[1:k+1], a[i])
		if p > 0 {
			ans = min(ans, i-p)
		}
	}

	if ans != 1<<60 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
