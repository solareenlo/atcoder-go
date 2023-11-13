package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)
	base := int(1e9)
	ans := 0
	for i := 10 - 1; i >= 0; i-- {
		y := min(x/(a*base+b), 9)
		x -= y * (a*base + b)
		ans += y * base
		base /= 10
	}
	ans = min(ans, int(1e9))
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
