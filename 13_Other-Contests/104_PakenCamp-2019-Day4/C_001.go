package main

import "fmt"

func main() {
	var A, B, C, D, E, F int
	fmt.Scan(&A, &B, &C, &D, &E, &F)
	ans := 0
	for i := A; i <= B; i++ {
		ans += max(0, min(D-i, D-C+1)) * max(0, min(F-i, F-E+1))
	}
	fmt.Println(float64(ans) / (float64(B-A+1) * float64(D-C+1) * float64(F-E+1)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
