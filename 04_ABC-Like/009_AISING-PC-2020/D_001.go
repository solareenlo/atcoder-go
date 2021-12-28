package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func f(x int) int {
	if x == 0 {
		return 0
	}
	return f(x%bits.OnesCount32(uint32(x))) + 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var X string
	fmt.Fscan(in, &n, &X)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = int(X[i] - '0')
	}

	pc := 0
	for i := 0; i < n; i++ {
		pc += x[i]
	}

	res := make([]int, n)
	for bit := 0; bit < 2; bit++ {
		next_pc := pc
		if bit == 0 {
			next_pc++
		} else {
			next_pc--
		}
		if next_pc <= 0 {
			continue
		}
		rem0 := 0
		for i := 0; i < n; i++ {
			rem0 = (rem0 * 2) % next_pc
			rem0 += x[i]
		}
		b := 1
		for i := n - 1; i >= 0; i-- {
			if x[i] == bit {
				rem := rem0
				if bit == 0 {
					rem = (rem + b) % next_pc
				} else {
					rem = (rem - b + next_pc) % next_pc
				}
				res[i] = f(rem) + 1
			}
			b = (b * 2) % next_pc
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, res[i])
	}
}
