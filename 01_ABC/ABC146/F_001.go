package main

import "fmt"

func main() {
	var n, m int
	var s string
	fmt.Scan(&n, &m, &s)

	res := make([]int, 0)
	j := 0
	for i := n; i > 0; i -= j {
		for j = m; j > 0; j-- {
			if i-j < 0 || s[i-j] == '1' {
				continue
			}
			res = append(res, j)
			break
		}
		if j == 0 {
			fmt.Println(-1)
			return
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
