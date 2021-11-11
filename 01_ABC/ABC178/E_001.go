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

	a := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[i] = x + y
		s[i] = x - y
	}
	sort.Ints(a)
	sort.Ints(s)

	fmt.Println(max(a[n-1]-a[0], s[n-1]-s[0]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
