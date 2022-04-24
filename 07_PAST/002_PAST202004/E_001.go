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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 1; i <= n; i++ {
		pre := a[i]
		cnt := 1
		for pre != i {
			pre = a[pre]
			cnt++
		}
		fmt.Fprint(out, cnt, " ")
	}
}
