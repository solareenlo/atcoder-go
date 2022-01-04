package main

import "fmt"

func main() {
	var n, m int
	var name, kit string
	fmt.Scan(&n, &m, &name, &kit)

	a := make([]int, 26)
	b := make([]int, 26)
	for i := 0; i < n; i++ {
		a[name[i]-'A']++
	}
	for i := 0; i < m; i++ {
		b[kit[i]-'A']++
	}

	res := 0
	for i := 0; i < 26; i++ {
		if a[i] != 0 {
			if b[i] == 0 {
				fmt.Println(-1)
				return
			}
			res = max(res, (a[i]+b[i]-1)/b[i])
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
