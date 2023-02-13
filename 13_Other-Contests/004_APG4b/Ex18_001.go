package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	date := make([][]byte, n)
	for i := range date {
		date[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			date[i][j] = '-'
		}
	}
	for i := 0; i < m; i++ {
		date[a[i]-1][b[i]-1] = 'o'
		date[b[i]-1][a[i]-1] = 'x'
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != n-1 {
				fmt.Printf("%c ", date[i][j])
			} else {
				fmt.Printf("%c", date[i][j])
			}
		}
		fmt.Println()
	}
}
