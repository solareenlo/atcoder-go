package main

import (
	"fmt"
	"sort"
)

func main() {
	for T := 0; T < 200; T++ {
		const n, m int = 30, 50
		var a [n][m]int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Scan(&a[i][j])
			}
		}

		c := make([]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				for k := 0; k < n; k++ {
					if a[i][j] == a[k][j] {
						c[i]++
					}
				}
			}
		}

		ord := make([]int, n)
		for i := 0; i < n; i++ {
			ord[i] = i
		}
		sort.Slice(ord, func(i, j int) bool { return c[ord[i]] > c[ord[j]] })

		ans := make([]int, m)
		for i := 0; i < m; i++ {
			x := [4]int{}
			for j := 0; j < 10; j++ {
				x[a[ord[j]][i]] += 5
			}
			for j := 10; j < 20; j++ {
				x[a[ord[j]][i]] += 1
			}
			for k := 0; k < 4; k++ {
				if x[ans[i]] < x[k] {
					ans[i] = k
				}
			}
		}

		for i := 0; i < m; i++ {
			if i+1 == m {
				fmt.Println(ans[i])
			} else {
				fmt.Printf("%d ", ans[i])
			}
		}
	}
}
