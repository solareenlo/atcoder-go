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

	var T int
	fmt.Fscan(in, &T)

	for k := 0; k < T; k++ {
		var n int
		fmt.Fscan(in, &n)
		cnt := map[int]int{}
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
			cnt[a[i]]++
		}
		flag := 0
		for i := 1; i <= n; i++ {
			flag |= (cnt[a[i]] & 1)
		}
		if n&1 != 0 || flag == 0 {
			fmt.Fprintln(out, "Second")
		} else {
			fmt.Fprintln(out, "First")
		}
	}
}
