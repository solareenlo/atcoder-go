package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1010
	const B = 53

	var f [N][N]int
	var a [N]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	f[0][0] = 1
	ans := 0
	for i := 0; i < B; i++ {
		for j := 0; j < N; j++ {
			if f[i][j] != 0 {
				for t := (k >> i) & 1; t <= ((k>>i)&1)*n; t++ {
					if (j+a[t])&1 != 0 && (n&1 != 0 || (k < 1<<(i+1))) {
						ans ^= 1 << i
					}
					f[i+1][(j+a[t])>>1] ^= 1
				}
			}
		}
	}
	fmt.Println(ans)
}
