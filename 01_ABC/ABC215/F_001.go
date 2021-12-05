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

	v := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i][0], &v[i][1])
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i][0] < v[j][0]
	})

	ok, ng := 0, 1000000001
	for ng-ok > 1 {
		md := (ok + ng) / 2
		que := make([][2]int, 0)
		able := false
		mini, maxi := 1000000001, 0
		for _, p := range v {
			for len(que) > 0 {
				if que[0][0] > p[0]-md {
					break
				}
				mini = min(mini, que[0][1])
				maxi = max(maxi, que[0][1])
				que = que[1:]
			}
			if mini <= p[1]-md || maxi >= p[1]+md {
				able = true
			}
			que = append(que, p)
		}
		if able {
			ok = md
		} else {
			ng = md
		}
	}
	fmt.Println(ok)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
