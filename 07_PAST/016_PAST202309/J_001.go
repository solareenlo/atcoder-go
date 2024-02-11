package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var N, K int
var A [200200]int

func check(k int) bool {
	return (A[K-1]-A[0])/k < N
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &K)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &A[i])
	}
	g := 0
	for i := 0; i < K-1; i++ {
		g = gcd(g, A[i+1]-A[i])
	}
	ans := make([]int, 0)
	for k := 1; k*k <= g; k++ {
		if g%k != 0 {
			continue
		}
		if check(k) {
			ans = append(ans, k)
		}
		if k != g/k && check(g/k) {
			ans = append(ans, g/k)
		}
	}
	sort.Ints(ans)
	fmt.Println(len(ans))
	for i := range ans {
		fmt.Println(ans[i])
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
