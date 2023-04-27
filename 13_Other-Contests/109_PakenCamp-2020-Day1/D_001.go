package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var ans int
	if K <= N {
		ans = n3(K)
	} else if K <= N*2 {
		ans = n3(K) - 3*n3(K-N)
	} else if K <= N*3 {
		ans = 6*n3(N) - n3(N*3-K)
	} else {
		ans = 6 * n3(N)
	}
	fmt.Println(ans)
}

func n3(n int) int {
	return n * n * n
}
