package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	ans := 0
	ac := 0
	for i := 0; i < n; i++ {
		ans += ac * (m - upperBound(b, a[i])) % mod
		ac += lowerBound(b, a[i])
	}
	fmt.Println(ans % mod)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
