package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	m := map[int]int{}
	for i := 0; i < l; i++ {
		var s int
		fmt.Scan(&s)
		m[s]++
	}

	cnt := 0
	for i := 0; i < r; i++ {
		var s int
		fmt.Scan(&s)
		if m[s] > 0 {
			cnt++
			m[s]--
		}
	}
	fmt.Println(cnt)
}
