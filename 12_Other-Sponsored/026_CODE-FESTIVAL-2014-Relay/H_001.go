package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, m)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	r := 0
	ans := 0
	for l := 0; l < m; l++ {
		for r < m && a[r] <= a[l]+n {
			r++
		}
		ans = max(ans, r-l)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
