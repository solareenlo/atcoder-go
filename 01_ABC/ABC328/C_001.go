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
	var s string
	fmt.Fscan(in, &n, &q, &s)
	cnt := make([]int, n)
	sum := 0
	for i := 0; i < n-1; i++ {
		cnt[i] = sum
		if s[i] == s[i+1] {
			sum++
		}
	}
	cnt[n-1] = sum
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		fmt.Fprintln(out, cnt[r]-cnt[l])
	}
}
