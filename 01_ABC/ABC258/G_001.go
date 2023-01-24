package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var v [3010][50]int
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < i; j++ {
			if s[j] == '1' {
				v[i][j>>6] |= 1 << (j & 63)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if (v[i][j>>6]>>(j&63))&1 != 0 {
				for k := 0; k < ((j + 63) >> 6); k++ {
					ans += bits.OnesCount(uint(v[i][k] & v[j][k]))
				}
			}
		}
	}
	fmt.Println(ans)
}
