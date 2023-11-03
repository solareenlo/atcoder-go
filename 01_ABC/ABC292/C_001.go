package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	num := make([]int, N+1)
	for i := 1; i*i <= N; i++ {
		for j := i; i*j <= N; j++ {
			num[i*j]++
			if i != j {
				num[i*j]++
			}
		}
	}

	ans := 0
	for i := 1; i < N; i++ {
		ans += num[i] * num[N-i]
	}
	fmt.Println(ans)
}
