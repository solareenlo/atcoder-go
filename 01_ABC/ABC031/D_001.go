package main

import (
	"fmt"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	v := make([]string, n)
	w := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	l := make([]int, 10)
	var dfs func(num int) bool
	dfs = func(num int) bool {
		if num == k+1 {
			for i := 0; i < n; i++ {
				sum := 0
				for _, u := range v[i] {
					sum += l[u-'0']
				}
				if sum != len(w[i]) {
					return false
				}
			}
			return true
		}
		for i := 1; i <= 3; i++ {
			l[num] = i
			if dfs(num + 1) {
				return true
			}
		}
		return false
	}
	dfs(1)

	res := make([]string, k+1)
	for i := 0; i < n; i++ {
		le := 0
		for _, u := range v[i] {
			num := u - '0'
			res[num] = w[i][le : le+l[num]]
			le += l[num]
		}
	}
	for i := 1; i <= k; i++ {
		fmt.Println(res[i])
	}
}
