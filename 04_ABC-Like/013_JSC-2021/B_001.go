package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	cnt := [1001]int{}
	for i := 0; i < n+m; i++ {
		var a int
		fmt.Scan(&a)
		cnt[a]++
	}

	for i := 0; i < 1001; i++ {
		if cnt[i] == 1 {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
