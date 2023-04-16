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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	for i := k; i < n; i++ {
		a[i] *= 2
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i]
	}
	fmt.Println(ans)
}
