package main

import "fmt"

func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)

	res := 1 << 60
	for i := 0; i < 202020; i++ {
		sum := c * i
		X := x - i/2
		Y := y - i/2
		if 0 < X {
			sum += X * a
		}
		if 0 < Y {
			sum += Y * b
		}
		res = min(res, sum)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
