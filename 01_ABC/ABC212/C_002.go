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
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	b := make([]int, m)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b)

	res := 1_000_000_000_000
	i, j := 0, 0
	for i < n && j < m {
		res = min(res, abs(a[i]-b[j]))
		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
