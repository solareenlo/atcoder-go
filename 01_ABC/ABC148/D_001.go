package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	num, cnt := 1, 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if num == tmp {
			num++
			cnt++
		}
	}

	if cnt == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(n - cnt)
	}
}
