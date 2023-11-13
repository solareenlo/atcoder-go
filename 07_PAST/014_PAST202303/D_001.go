package main

import "fmt"

func main() {
	var H, A, B, C, D int
	fmt.Scan(&H, &A, &B, &C, &D)
	ans := (H + A - 1) / A * B
	t := 0
	for H > 0 {
		t += D
		H -= C
		if H > 0 {
			H -= H / 2
		}
		ans = min(ans, t+(max(0, H)+A-1)/A*B)
	}
	fmt.Println(ans)
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
