package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	W := make([]int, 1<<13)
	D := make([]int, 1<<13)
	for {
		var n, a, b int
		fmt.Fscan(in, &n, &a, &b)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &W[1<<i])
		}
		for i := 1; i < 1<<n; i++ {
			W[i] = W[i&-i] + W[i-(i&-i)]
		}
		for i := range D {
			D[i] = 0
		}
		D[0] = 1
		for i := 0; i < 1<<n; i++ {
			for j := i; ; j = (j - 1) & i {
				k := i ^ j
				s := abs(W[j] - W[k])
				if i == 0 || a <= s && s <= b {
					for p := 0; p < n; p++ {
						if ((i >> p) & 1) == 0 {
							D[i|1<<p] += D[j] * D[k]
						}
					}
				}
				if j == 0 {
					break
				}
			}
		}
		fmt.Fprintln(out, D[(1<<n)-1])
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
