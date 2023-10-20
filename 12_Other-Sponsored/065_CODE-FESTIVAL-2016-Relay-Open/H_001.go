package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const m = 10800
	const d = 86400

	var n int
	fmt.Fscan(in, &n)
	var a, b [100000]int
	v := make([]int, 0)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		v = append(v, sum+a[i])
		sum += (a[i] + b[i])
	}
	for i := 0; i < n; i++ {
		v[i] %= d
		v = append(v, v[i]+d)
	}
	sort.Ints(v)
	sax := 0
	for i := 0; i < n; i++ {
		p := upperBound(v, v[i]+m)
		if (p - i) > sax {
			sax = p - i
		}
	}
	fmt.Println(sax)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
