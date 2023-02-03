package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)

	maxN, k, ans := 0, 1, 0
	for T > 0 {
		T--
		var x int
		fmt.Fscan(in, &x)
		if x > maxN {
			maxN = x
			ans = k
		}
		k++
	}
	fmt.Println(ans)
}
