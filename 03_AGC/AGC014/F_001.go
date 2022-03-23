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

	p := make([]int, 200001)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		p[x] = i
	}

	ans := 0
	g := make([]int, 200001)
	for i := n - 1; i >= 1; i-- {
		tmp1, tmp2, tmp3 := 0, 0, 0
		if p[i] < p[i+1] {
			tmp1 = 1
		}
		if p[i+1] < p[g[i+1]] {
			tmp2 = 1
		}
		if p[g[i+1]] < p[i] {
			tmp3 = 1
		}
		if ans == 0 {
			if p[i] > p[i+1] {
				ans = 1
				g[i] = i + 1
			}
		} else if tmp1+tmp2+tmp3 == 2 {
			g[i] = g[i+1]
		} else {
			ans++
			g[i] = i + 1
		}
	}
	fmt.Println(ans)
}
