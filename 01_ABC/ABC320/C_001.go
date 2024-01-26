package main

import "fmt"

func main() {
	var n int
	var s1, s2, s3 string
	fmt.Scan(&n, &s1, &s2, &s3)
	m := n * 3
	ans := m
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if s1[i%n] == s2[j%n] && s1[i%n] == s3[k%n] && i != j && i != k && j != k {
					ans = min(ans, max(i, max(j, k)))
				}
			}
		}
	}
	if ans == m {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
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
