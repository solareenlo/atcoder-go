package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n int
	fmt.Fscan(in, &n)
	var f [1 << 22]int
	f[0] = 1
	var a [22][22]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for k := 1; k < 1<<n; k++ {
		i := bits.OnesCount(uint(k)) - 1
		for j := 0; j < n; j++ {
			if a[i][j] != 0 && (k&(1<<j)) != 0 {
				f[k] = (f[k] + f[k-(1<<j)]) % mod
			}
		}
	}
	fmt.Println(f[(1<<n)-1])
}
