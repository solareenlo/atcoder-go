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

	var n, m int
	fmt.Scan(&n, &m)

	type pair struct{ x, y int }
	s := make(map[pair]bool)
	for i := 1; i <= max(n, m); i += 2 {
		for j := min(i, n); j <= min(i+2, n); j++ {
			for k := min(i, m); k <= min(i+2, m); k++ {
				s[pair{j, k}] = true
			}
		}
	}

	keys := make([]pair, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].x < keys[j].x || (keys[i].x == keys[j].x && keys[i].y < keys[j].y)
	})

	fmt.Fprintln(out, len(s))
	for i := range keys {
		fmt.Fprintln(out, keys[i].x, keys[i].y)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
