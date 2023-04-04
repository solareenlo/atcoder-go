package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	var p, q int
	fmt.Fscan(in, &p, &q)
	pro := (mod + 1 - (p*powMod(q, mod-2))%mod) % mod
	sum, X := 0, 0
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		sum += A[i]
		X ^= A[i]
	}
	if X == 0 {
		fmt.Println(0)
		return
	}
	sort.Ints(A)
	if A[N-1] == 1 {
		fmt.Println(1)
		return
	}
	if N%2 == 1 && (N == 1 || A[N-2] == 1) {
		fmt.Println(pro)
		return
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans = max(ans, A[i]-(A[i]^X))
	}
	ans = (sum - ans) / 2
	if ans == 0 {
		ans++
	}
	fmt.Println(powMod(pro, ans))
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
