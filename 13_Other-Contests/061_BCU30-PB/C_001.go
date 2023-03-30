package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var used [300][300]bool
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		used[a][b] = true
	}

	res := 0
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			cnt := 0
			for m := 0; m < n; m++ {
				if used[a][m] && used[m][b] {
					cnt++
				}
			}
			res += cnt * cnt
		}
	}

	fmt.Println(res)
}
