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

	var H, W, N, MOD int
	fmt.Fscan(in, &H, &W, &N, &MOD)
	var res, p int = 0, 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var C int
			fmt.Fprintf(out, "? %d %d %d %d\n", i, j, i, j)
			out.Flush()
			fmt.Fscan(in, &C)
			res = (res + p*C%MOD) % MOD
			p = p * 2 % MOD
		}
	}
	fmt.Fprintf(out, "! %d\n", res)
	out.Flush()
}
