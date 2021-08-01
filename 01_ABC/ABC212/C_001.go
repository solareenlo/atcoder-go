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

	res := int(1e12)
	i := 0
	j := 0
	for i < n && j < m {
		if c := abs(a[i] - b[j]); c < res {
			res = c
		}
		if a[i] < b[j] {
			i++
		} else {
			j++
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
