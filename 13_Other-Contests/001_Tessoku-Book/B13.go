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
	B := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		B[i] = B[i-1] + a
	}
	ans := 0
	for i := 1; i <= n; i++ {
		x := 1
		for i-x >= 0 && B[i]-B[i-x] <= k {
			x++
		}
		ans += (x - 1)
	}
	fmt.Println(ans)
}
