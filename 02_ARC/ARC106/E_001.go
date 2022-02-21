package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	p := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		m[p[i]]++
	}

	o := 0
	c := 0
	for {
		s := 0
		for i := 0; i < n; i++ {
			if o%(p[i]*2) < p[i] {
				s = 1
				break
			}
		}
		c += s
		if c >= k*n {
			break
		}
		o++
	}
	o++

	for a, b := range m {
		b *= k
		x := (b / a) * a * 2
		if b%a != 0 {
			x += b % a
		} else {
			x -= a
		}
		if o < x {
			o = x
		}
	}

	fmt.Println(o)
}
