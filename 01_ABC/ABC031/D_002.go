package main

import "fmt"

func main() {
	var K, n int
	fmt.Scan(&K, &n)
	v := make([]string, n)
	w := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	res := make([]string, K+1)
	// v[i] の j 文字目，w[i] の k 文字目を見ている
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i == n {
			return true
		}
		if j == len(v[i]) {
			if k == len(w[i]) {
				return dfs(i+1, 0, 0)
			} else {
				return false
			}
		}
		d := v[i][j] - '0'
		l := len(res[d])
		if l != 0 {
			if res[d] == w[i][k:min(k+l, len(w[i]))] {
				return dfs(i, j+1, k+l)
			}
		} else {
			for t := 0; t < 3; t++ {
				if k+t < len(w[i]) {
					res[d] = w[i][k : k+t+1]
					if dfs(i, j+1, k+t+1) {
						return true
					}
				}
			}
			res[d] = ""
		}
		return false
	}
	dfs(0, 0, 0)
	for i := 0; i < K; i++ {
		fmt.Println(res[i+1])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
