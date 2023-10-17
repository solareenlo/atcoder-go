package main

import "fmt"

func main() {
	var n, m, a int
	fmt.Scan(&n, &m, &a)
	ax := (m + 1) / 2
	ay := (n + 1) / 2
	if ay*ax < a {
		fmt.Println("IMPOSSIBLE")
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i|j)&1 != 0 || a == 0 {
				fmt.Print(".")
			} else {
				a--
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
