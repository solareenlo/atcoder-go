package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f, s, res := 2, 1, 3
	for i := 0; i < n-1; i++ {
		res = f + s
		f = s
		s = res
	}

	if n == 0 {
		res = 2
	} else if n == 1 {
		res = 1
	}
	fmt.Println(res)
}
