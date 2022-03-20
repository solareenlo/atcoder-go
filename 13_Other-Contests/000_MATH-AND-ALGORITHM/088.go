package main

import "fmt"

const mod = 998244353

func f(X int) int {
	return X * (X + 1) / 2 % mod
}

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	fmt.Println(f(A) * f(B) % mod * f(C) % mod)
}
