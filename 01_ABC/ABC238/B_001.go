package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
		a[i] = (a[i] + a[i-1]) % 360
	}
	sort.Ints(a)

	maxi := a[1]
	a[n+1] = 360
	for i := 1; i <= n; i++ {
		maxi = max(maxi, a[i+1]-a[i])
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
