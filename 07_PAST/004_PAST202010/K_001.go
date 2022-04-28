package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var K int
	fmt.Fscan(in, &K)
	num := make([][20]int, K)
	internal := make([]int, K)
	for i := 0; i < K; i++ {
		var N int
		fmt.Fscan(in, &N)
		for j := 0; j < N; j++ {
			var a int
			fmt.Fscan(in, &a)
			a--
			num[i][a]++
			for k := 0; k < 20; k++ {
				if k > a {
					internal[i] += num[i][k]
				}
			}
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	ans := 0
	const mod = 1_000_000_000
	cnt := make([]int, 20)
	for q := 0; q < Q; q++ {
		var b int
		fmt.Fscan(in, &b)
		for i := 0; i < 20; i++ {
			for j := 0; j < 20; j++ {
				if i < j {
					ans = (ans + (num[b-1][i]*cnt[j])%mod) % mod
				}
			}
		}
		ans = (ans + internal[b-1]) % mod
		for i := 0; i < 20; i++ {
			cnt[i] = (cnt[i] + num[b-1][i]) % mod
		}
	}
	fmt.Println(ans)
}
