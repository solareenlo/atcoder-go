package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
	}

	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[a[i]]++
	}

	tot := 0
	for i := 0; i < n; i++ {
		tot += nC2(cnt[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(tot - (cnt[a[i]] - 1))
	}
}

func nC2(n int) int {
	return n * (n - 1) / 2
}
