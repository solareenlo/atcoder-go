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
	ans := 9
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		tmp := 0
		for p%10 == 0 {
			p /= 10
			tmp++
		}
		ans = min(ans, tmp)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
