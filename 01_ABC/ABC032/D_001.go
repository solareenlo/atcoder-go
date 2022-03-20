package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, W int
	fmt.Scan(&N, &W)

	v := make([]int, N)
	w := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	dp := make(map[int]int)
	dp[0] = 0
	for i := 0; i < N; i++ {
		tmp := make(map[int]int)
		for key, val := range dp {
			tmp[key] = val
		}
		keys := make([]int, 0, len(dp))
		for k := range dp {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for k := range keys {
			nw := keys[k] + w[i]
			if nw > W {
				continue
			}
			nv := dp[keys[k]] + v[i]
			if tmp[nw] < nv {
				tmp[nw] = nv
			}
		}
		keys = make([]int, 0, len(tmp))
		for k := range tmp {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		next_dp := make(map[int]int)
		current := -1
		for k := range keys {
			if tmp[keys[k]] > current {
				current = tmp[keys[k]]
				next_dp[keys[k]] = tmp[keys[k]]
			}
		}
		dp = next_dp
	}

	keys := make([]int, 0, len(dp))
	for k := range dp {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println(dp[keys[len(keys)-1]])
}
