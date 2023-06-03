package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if K-i-j > 0 && K-i-j <= N {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}
