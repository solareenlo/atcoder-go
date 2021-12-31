package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	ok := make([]bool, N)
	for i := range ok {
		ok[i] = true
	}
	for i := 0; i < M; i++ {
		var a int
		fmt.Fscan(in, &a)
		ok[a] = false
	}

	base := make([]int, 0)
	elim := make([]int, 0)
	for x := 1; x < N; x++ {
		if ok[x] {
			y := x
			for _, b := range elim {
				y = min(y, y^b)
			}
			if y != 0 {
				base = append(base, x)
				elim = append(elim, y)
			}
		}
	}

	__lg := func(x int) int {
		b := x & -x
		return bits.OnesCount(uint(b - 1))
	}

	if len(base) != __lg(N) {
		fmt.Fprintln(out, -1)
		return
	}

	XOR := 0
	for x := 1; x < N; x++ {
		fmt.Fprint(out, XOR, " ")
		XOR ^= base[__lg(x&-x)]
		fmt.Fprintln(out, XOR)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
