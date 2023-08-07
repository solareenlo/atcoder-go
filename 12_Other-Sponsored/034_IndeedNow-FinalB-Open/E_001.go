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
	var a [100009]int
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		p[i] = i
	}
	sort.Slice(p, func(i, j int) bool {
		return a[p[i]] > a[p[j]]
	})
	ret := 0
	ok := true
	var bit [100009]int
	for i := 0; i < n; i++ {
		for j := p[i]; j >= 1; j -= j & (-j) {
			ret += bit[j] * a[p[i]]
		}
		for j := p[i] + 1; j <= n; j += j & (-j) {
			bit[j]++
		}
		if i != 0 && a[p[i]] == a[p[i-1]] {
			ok = false
		}
	}
	if ok {
		fmt.Println(ret)
	} else {
		fmt.Println(-1)
	}
}
