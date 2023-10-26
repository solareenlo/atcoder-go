package main

import (
	"fmt"
)

func main() {
	var N, K, R int
	fmt.Scan(&N, &K, &R)
	ans := make([]int, N+1)
	l, r := 1, N
	for i := 1; i <= N; i++ {
		if N-(i+K-1) >= 0 && R-(N-(i+K-1)) >= 0 {
			ans[r] = i
			r--
			R -= N - (i + K - 1)
		} else {
			ans[l] = i
			l++
		}
	}
	if R != 0 {
		fmt.Println("No Luck")
		return
	}
	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", ans[i])
	}
}
