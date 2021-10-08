package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	sumA := make([]int, n+1)
	sumB := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		sumA[i] = sumA[i-1] + a[i]
		sumB[i] = sumB[i-1] ^ a[i]
	}

	i, res := 1, 0
	for j := 0; j < n+1; j++ {
		for (sumA[j] - sumA[i-1]) != (sumB[j] ^ sumB[i-1]) {
			i++
		}
		res += j - i + 1
	}
	fmt.Println(res)
}
