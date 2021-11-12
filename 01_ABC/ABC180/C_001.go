package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	div := divList(n)

	keys := make([]int, 0, len(div))
	for k := range div {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Fprintln(out, k)
	}
}

func divList(n int) map[int]struct{} {
	div := map[int]struct{}{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			div[i] = struct{}{}
			if i*i != n {
				div[n/i] = struct{}{}
			}
		}
	}
	return div
}
