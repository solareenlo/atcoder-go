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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	res := a[n-1] - a[0]
	for i := 1; i < n/2; i++ {
		res += a[n-i] - a[i]
		res += a[n-i-1] - a[i-1]
	}

	if n&1 != 0 {
		res += max(a[n/2+1]-a[n/2], a[n/2]-a[n/2-1])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
