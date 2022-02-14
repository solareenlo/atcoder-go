package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	ans := 0
	if n&1 != 0 {
		for i := 0; i < n; i++ {
			ans ^= a[i]
		}
		for i := 0; i < n; i++ {
			ans ^= b[i]
		}
	}

	for i := 0; i < n; i++ {
		a[i] = ^a[i]
	}

	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < 28; i++ {
		p := n
		q := n
		s := 0
		t := 0
		mask := (2 << i) - 1
		for j := 0; j < n; j++ {
			p -= (a[j] >> i) & 1
		}
		for j := 0; j < n; j++ {
			q -= (b[j] >> i) & 1
		}
		for j := 0; j < n; j++ {
			if (a[j]>>i)&1 != 0 {
				c[p] = a[j]
				p++
			} else {
				c[s] = a[j]
				s++
			}
		}
		for j := 0; j < n; j++ {
			if (b[j]>>i)&1 != 0 {
				d[q] = b[j]
				q++
			} else {
				d[t] = b[j]
				t++
			}
		}
		a, c = c, a
		b, d = d, b
		t = 0
		for s := 0; s < n; s++ {
			for t < n && (a[s]&mask) >= (b[t]&mask) {
				t++
			}
			if (n-t)&1 != 0 {
				ans ^= 2 << i
			}
		}
	}
	fmt.Println(ans)
}
