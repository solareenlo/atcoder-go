package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	ans := 0
	for i := 0; i < n; i++ {
		var I int
		fmt.Fscan(in, &I)
		if I == x {
			ans = i
			break
		}
	}
	fmt.Println(ans + 1)
}
