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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := int(1e18)
	for i := 0; i < n; i++ {
		tmp := divCnt(&a[i], 2)
		ans = min(ans, tmp)
	}
	fmt.Println(ans)
}

func divCnt(num *int, div int) int {
	cnt := 0
	for *num%div == 0 {
		*num /= div
		cnt++
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
