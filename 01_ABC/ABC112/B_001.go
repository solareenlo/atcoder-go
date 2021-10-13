package main

import "fmt"

func main() {
	var N, T, c, t int
	fmt.Scan(&N, &T)

	res := 1001
	for i := 0; i < N; i++ {
		fmt.Scan(&c, &t)
		if T >= t {
			res = min(res, c)
		}
	}
	if res == 1001 {
		fmt.Println("TLE")
	} else {
		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
