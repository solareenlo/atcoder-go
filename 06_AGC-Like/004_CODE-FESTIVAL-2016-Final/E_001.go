package main

import (
	"fmt"
	"math"
)

func main() {
	var N, A int
	fmt.Scan(&N, &A)

	ans := N
	for k := 1; k < 40; k++ {
		p := int(math.Pow(float64(N), 1.0/float64(k+1)))
		now := 1
		for i := 0; i < k; i++ {
			now *= p
		}
		for i := 0; i <= k; i++ {
			ans = min(ans, A*k+p*k+i+(N+now-1)/now)
			now = now / p * (p + 1)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
