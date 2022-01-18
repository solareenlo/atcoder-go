package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := map[int]bool{}
	for i := 0; i < n+m; i++ {
		var t int
		fmt.Scan(&t)
		s[t] = true
	}

	r := len(s)
	fmt.Println(float64(n+m-r) / float64(r))
}
