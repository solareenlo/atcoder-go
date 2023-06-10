package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007
	var A [55555]int

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N+100; i++ {
		if i < N {
			var a int
			fmt.Fscan(in, &a)
			A[i] += a
		}
		t := (A[i] - 9) / 10
		A[i] -= t * 10
		A[i+1] += t
	}
	ans := 1
	pre := 0
	for i := N + 100 - 1; i >= 0; i-- {
		if A[i] <= 8 {
			pre = ans
			ans *= A[i] + 1
		} else {
			ans = ans*10 + pre*(A[i]-9)
		}
		ans %= mod
	}
	fmt.Println((ans + mod - 1) % mod)
}
