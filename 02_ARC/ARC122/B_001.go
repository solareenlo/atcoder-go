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

	num := a[n/2]
	res := 0
	for i := 0; i < n; i++ {
		res += a[i] - min(a[i], num)
	}
	fmt.Println(float64(res)/float64(n) + float64(num)/2.0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
