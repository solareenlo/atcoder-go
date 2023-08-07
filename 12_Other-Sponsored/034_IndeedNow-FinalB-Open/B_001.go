package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200007

	var n int
	fmt.Fscan(in, &n)
	var s, t, k, pre [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i], &t[i])
		k[s[i]]++
	}
	for i := 1; i <= 2*n; i++ {
		pre[i] = pre[i-1] + k[i]
	}
	for i := 1; i <= n; i++ {
		fmt.Println(pre[t[i]] - pre[s[i]-1] - 1)
	}
}
