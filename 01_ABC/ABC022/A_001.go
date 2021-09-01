package main

import "fmt"

func main() {
	var n, s, t, w, a int
	fmt.Scan(&n, &s, &t)

	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		w += a
		if s <= w && w <= t {
			res++
		}
	}
	fmt.Println(res)
}
