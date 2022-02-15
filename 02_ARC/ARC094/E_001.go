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

	ans := 0
	m := 1 << 60
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		ans += a
		if a > b && m > b {
			m = b
		}
	}

	if m > 1_000_000_000 {
		fmt.Println(0)
	} else {
		fmt.Println(ans - m)
	}
}
