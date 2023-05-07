package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	A := make([]int, K)
	for i := 0; i < K; i++ {
		if i < 9 {
			A[i] = powMod(2, i)
		}
		if i == 9 {
			A[i] = 511
		}
		if i > 9 {
			A[i] = (2 * A[i-1] % 1000000007) - A[i-10]
			for A[i] < 0 {
				A[i] += 1000000007
			}
		}
	}
	if K%9 != 0 {
		fmt.Println(0)
	} else {
		fmt.Println(A[K-1])
	}
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
