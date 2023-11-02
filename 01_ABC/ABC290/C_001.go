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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	v := make([]int, 0)
	v = append(v, a[1])
	for i := 2; i <= n; i++ {
		if a[i] != a[i-1] {
			v = append(v, a[i])
		}
	}
	now := 0
	m := len(v)
	for i := 0; i < min(k, m); i++ {
		if v[i] == now {
			now++
		} else {
			break
		}
	}
	fmt.Println(now)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
