package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	x := make([]int, 30)
	for j := 0; j < t; j++ {
		var a int
		fmt.Fscan(in, &a)
		ans := 0
		n := 0
		for a > 0 {
			x[n] = a % 10
			a /= 10
			n++
		}
		for i := n - 1; i >= 0; i-- {
			ans = max(ans, (x[i]+2)/3)
			tmp := 0
			if i != 0 && (x[i-1] < ans) {
				tmp = 1
			}
			if x[i]-tmp < ans {
				ans = max(ans, (x[i]+12-tmp)/3)
			}
		}
		fmt.Println(ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
