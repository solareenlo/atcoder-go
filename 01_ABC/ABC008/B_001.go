package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	m := map[string]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		m[s]++
	}

	x := 0
	for k, v := range m {
		if v > x {
			x = v
			s = k
		}
	}
	fmt.Println(s)
}
