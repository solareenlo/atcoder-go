package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	ans := map[int]int{}
	cur := map[int]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		nxt := map[int]int{}
		nxt[a]++
		for k, v := range cur {
			nxt[gcd(k, a)] += v
		}
		for k, v := range nxt {
			ans[k] += v
		}
		cur = nxt
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		fmt.Fprintln(out, ans[x])
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
