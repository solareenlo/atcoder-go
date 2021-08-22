package main

import "fmt"

var r, g, b int

func cnt(p int) int {
	le := min(p-r, -100-(r/2))
	re := max(p+g, 100-(b/2))
	res := 0
	for i := p + g - 1; i >= p; i-- {
		res += abs(i)
	}
	for i := le + r - 1; i >= le; i-- {
		res += abs(-100 - i)
	}
	for i := re + b - 1; i >= re; i-- {
		res += abs(100 - i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Scan(&r, &g, &b)
	res := int(1e9 + 7)
	for i := -300; i <= 300; i++ {
		res = min(res, cnt(i))
	}
	fmt.Println(res)
}
