package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	ans := 0
	for i, j := 0, 1; j < len(s); j++ {
		if s[j] != s[i] {
			ans += (j - i) * (n - j)
			i = j
		}
	}
	fmt.Println(ans)
}
