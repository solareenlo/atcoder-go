package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	m := map[int]bool{}
	for i := 0; i < n; i++ {
		m[i] = false
	}

	for i := 0; i < k; i++ {
		var tmp int
		fmt.Scan(&tmp)
		for j := 0; j < tmp; j++ {
			var tmp2 int
			fmt.Scan(&tmp2)
			m[tmp2-1] = true
		}
	}

	cnt := 0
	for _, v := range m {
		if !v {
			cnt++
		}
	}
	fmt.Println(cnt)
}
