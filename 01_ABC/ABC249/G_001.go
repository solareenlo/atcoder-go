package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	b := make([]int, 0)
	d := make([]int, 0)
	v := K + 1
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		y = (y << 30) | x
		for _, z := range b {
			y = min(y, y^z)
		}
		for _, z := range d {
			x = min(x, x^z)
		}
		v = min(v, x)
		if x != 0 {
			d = append(d, x)
		}
		if y != 0 {
			b = append(b, y)
		}
	}
	if v > K {
		fmt.Println(-1)
		return
	}

	B := (1 << 30) - 1
	ret := 0
	for i := 0; i < len(b); i++ {
		ret = max(ret, ret^b[i])
		c := make([]int, 0)
		for j := i + 1; j < len(b); j++ {
			x := b[j] & B
			for _, z := range c {
				x = min(x, z^x)
			}
			if x != 0 {
				c = append(c, x)
			}
		}
		f := ret & B
		for _, z := range c {
			f = min(f, f^z)
		}
		if f > K {
			ret ^= b[i]
		}
	}
	fmt.Println(ret >> 30)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
