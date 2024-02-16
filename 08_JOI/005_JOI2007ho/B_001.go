package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var c [1 << 17]int
	A := 0
	for i := 0; i < k; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a > 0 {
			c[a] = 1
		} else {
			A = 1
		}
	}
	ans := 0
	f := 1
	cnt := 0
	for i := 1; i <= n; i++ {
		if c[i] < 1 {
			cnt++
		}
		for f <= i && cnt > A {
			if c[f] < 1 {
				cnt--
			}
			f++
		}
		if ans < i-f+1 {
			ans = i - f + 1
		}
	}
	fmt.Println(ans)
}
