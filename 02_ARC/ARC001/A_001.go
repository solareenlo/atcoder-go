package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	mp := map[byte]int{}
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}

	keys := make([]byte, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	maxi := 0
	mini := 101
	if len(mp) == 1 {
		maxi = mp[keys[0]]
		mini = 0
	} else {
		for _, v := range mp {
			maxi = max(maxi, v)
			mini = min(mini, v)
		}
	}

	fmt.Println(maxi, mini)
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
