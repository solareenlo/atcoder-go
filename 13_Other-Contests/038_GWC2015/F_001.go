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

	s := 0
	m := 55
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		s += x
		m = min(m, x)
	}

	if ((n&1) != 0 && (s&1) != 0) || (n&1) == 0 && ((m&1) != 0 || (s&1) != 0) {
		fmt.Println("Iori")
	} else {
		fmt.Println("Yayoi")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
