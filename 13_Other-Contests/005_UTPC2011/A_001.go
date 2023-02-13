package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res := 0
	for n > 0 {
		n--
		cnt := 0
		for i := 0; i < m; i++ {
			var a int
			fmt.Scan(&a)
			if a != 0 {
				cnt++
			}
		}
		res = max(res, cnt)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
