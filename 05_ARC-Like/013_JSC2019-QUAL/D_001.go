package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Fprint(out, ctz(i^j)+1, " ")
		}
		fmt.Fprintln(out)
	}
}

func ctz(x int) int {
	return bits.TrailingZeros64(uint64(x))
}
