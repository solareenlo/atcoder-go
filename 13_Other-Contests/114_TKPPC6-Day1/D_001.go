package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var N int
	fmt.Fscan(in, &N)

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}

	acc := make([]int, N+1)
	for i := 0; i < N; i++ {
		acc[i+1] += acc[i] + a[i]
	}
	sort.Ints(acc)

	ans := 0
	for i := 0; i < N+1; i++ {
		ans = (ans + (acc[i]+MOD)%MOD*(i*2-N) + MOD) % MOD
	}

	fmt.Println(ans)
}
