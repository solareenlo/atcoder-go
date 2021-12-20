package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(in, &n, &w)

	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	res := 0
	for i := n - 1; i >= 0; i-- {
		t := min(w, a[i][1])
		w -= t
		res += t * a[i][0]
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
