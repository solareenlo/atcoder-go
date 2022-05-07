package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][1])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	st := make(map[int]bool)
	ans := 0
	for i := 0; i < n; i++ {
		if len(st) < k {
			if _, ok := st[a[i][1]]; ok {
				continue
			}
			st[a[i][1]] = true
			ans += a[i][0]
		}
	}

	if len(st) < k {
		ans = -1
	}
	fmt.Println(ans)
}
