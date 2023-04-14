package main

import "fmt"

func main() {
	var A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z int
	fmt.Scan(&A, &B, &C, &D, &E, &F, &G, &H, &I, &J, &K, &L, &M, &N, &O, &P, &Q, &R, &S, &T, &U, &V, &W, &X, &Y, &Z)
	s := []string{"a", "aa", "aaa", "aaai", "aaaji", "aabaji", "agabaji", "dagabaji"}
	a := []int{6, 28, 496, 8128}
	var i, j int
	for k := 1; k < 3600; k++ {
		if k%59 == K && k%61 == L {
			i = k + 3599*(M-1)
		}
	}
	for k := 0; k < 4; k++ {
		if abs(i-a[k]) >= N {
			j = a[k]
			break
		}
	}
	fmt.Println(A - B)
	fmt.Println(C + D)
	fmt.Println(max(F-E+1, 0))
	fmt.Println((G + H + I + 3) / 3)
	fmt.Println(s[J-1])
	fmt.Println(min(i, j))
	fmt.Println(max(i, j))
	fmt.Println((O + P + Q) * (R + S + T) * (U + V + W) * (X + Y + Z) % 9973)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
