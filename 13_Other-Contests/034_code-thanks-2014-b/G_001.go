package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p int
	fmt.Fscan(in, &n, &p)
	m := n % 5
	if (m == 2 && p == 1) || (m == 3 && p <= 2) {
		fmt.Println("second")
	} else if m == 0 {
		q := n / 5
		cnt := 5 << (ctz(q))
		if p <= cnt-1 {
			fmt.Println("second")
		} else {
			fmt.Println("first")
		}
	} else {
		fmt.Println("first")
	}
}

func ctz(x int) int {
	return bits.TrailingZeros32(uint32(x))
}
