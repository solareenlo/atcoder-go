package main

import "fmt"

func main() {
	var n, t, d, m, res int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t, &d, &m)
		if t+10 <= d {
			res += m
		}
	}
	fmt.Println(res)
}
