package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N, A, B, X, Y, Z int
		fmt.Scan(&N, &A, &B, &X, &Y, &Z)
		Y = min(Y, A*X)
		Z = min(Z, B*X)
		if A < B {
			A, B = B, A
			Y, Z = Z, Y
		}
		ans := 1 << 60
		for k1 := 0; k1 <= min(N/A, B-1); k1++ {
			for _, k2 := range [2]int{0, (N/A - k1) / B} {
				k := k1 + k2*B
				l := (N - k*A) / B
				ans = min(ans, N*X+k*(Y-A*X)+l*(Z-B*X))
			}
		}
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
