package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 400400

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, N)
	als := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		als += a[i]
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	for i := 1; i <= n; i++ {
		a[i+n] = a[i] + m
	}

	i := 1
	ans := int(1e18)
	for i <= n {
		p := i + 1
		sum := a[i]
		for a[p]-a[p-1] <= 1 && p-i < n {
			sum += a[p] % m
			p++
		}
		i = p
		ans = min(als-sum, ans)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
