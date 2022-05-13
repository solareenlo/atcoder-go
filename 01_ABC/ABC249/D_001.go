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

	b := make([]int, 200005)
	maxn := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		b[a]++
		if a > maxn {
			maxn = a
		}
	}

	ans := 0
	for i := 1; i <= maxn; i++ {
		for j := 1; i*j <= maxn; j++ {
			ans = ans + b[i]*b[j]*b[i*j]
		}
	}
	fmt.Println(ans)
}
