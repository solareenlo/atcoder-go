package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [300030]int
	var op [30]int

	var n int
	fmt.Fscan(in, &n)
	S := (1 << n) - 1
	for i := 1; i <= S; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		a[i] = tmp ^ a[i-1]
		if a[i] == 0 {
			continue
		}
		op[32-countLeadingZeros(uint32(a[i]))] = 1
	}
	ans := 1
	for i := 1; i <= n; i++ {
		if op[i] == 0 {
			ans <<= 1
		}
	}
	fmt.Println(ans)
}

func countLeadingZeros(x uint32) int {
	return bits.LeadingZeros32(x)
}
