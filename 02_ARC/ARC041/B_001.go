package main

import "fmt"

var (
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{-1, 0, 1, 0}
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < m; j++ {
			b[i][j] = int(s[j] - '0')
		}
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := 0; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			res[i+1][j] = b[i][j]
			tmp := b[i][j]
			for k := 0; k < 4; k++ {
				if i+1+dy[k] >= 0 && i+1+dy[k] < n {
					if j+dx[k] >= 0 && j+dx[k] < m {
						b[i+1+dy[k]][j+dx[k]] -= tmp
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(res[i][j])
		}
		fmt.Println()
	}
}
