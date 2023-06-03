package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		L[i] = 1
		R[i] = 1
	}
	for i := 0; i < N-1; i++ {
		if S[i] == 'A' {
			L[i+1] = L[i] + 1
		}
		if S[N-2-i] == 'B' {
			R[N-2-i] = R[N-1-i] + 1
		}
	}
	ans := 0
	for N > 0 {
		N--
		ans += max(L[N], R[N])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
