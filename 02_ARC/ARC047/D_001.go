package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1000000007
const N = 5007

var (
	n  int
	q  int
	s  = [2 * N]int{}
	t  = [2 * N]int{}
	r1 int
	r2 int
)

func update(x, c int, s, b *int) {
	if *s < x {
		*b = c
		*s = x
	} else if *s == x {
		*b += c
	}
}

func query(x, y, l int) {
	sb := x + y + (l - 1)
	tb := x - y + n
	if l == 1 {
		r1 = s[sb] + t[tb]
		r2 = 1
		return
	}
	ma := s[sb]
	b := 1
	c := 1
	if t[tb+(1-l)] == t[tb+(l-1)] {
		c++
	}
	r := ma + max(t[tb+(1-l)], t[tb+(l-1)])
	for k := 2; k < l; k += 2 {
		update(s[sb+k], 1, &ma, &b)
		update(s[sb-k], 1, &ma, &b)
		update(ma+t[tb+(l-1)-k], b, &r, &c)
		if k < l-1 {
			update(ma+t[tb-(l-1)+k], b, &r, &c)
		}
	}
	ma = -INF
	b = 0
	for k := 1; k < l; k += 2 {
		update(s[sb+k], 1, &ma, &b)
		update(s[sb-k], 1, &ma, &b)

		update(ma+t[tb+(l-1)-k], b, &r, &c)
		if k < l-1 {
			update(ma+t[tb-(l-1)+k], b, &r, &c)
		}
	}
	r1 = r
	r2 = c
}

func solve(a, b, c, d int) {
	b++
	d++
	res := -INF
	cnt := 0
	for a < b && c < d {
		if b-a < d-c {
			l := b - a
			query(a, c, l)
			update(r1, r2, &res, &cnt)
			c += l
		} else {
			l := d - c
			query(a, c, l)
			update(r1, r2, &res, &cnt)
			a += l
		}
	}
	r1 = res
	r2 = cnt
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n, &q)
	for i := 0; i < q; i++ {
		var tp int
		fmt.Fscan(in, &tp)
		var a, b, c, d int
		switch tp {
		case 1:
			fmt.Fscan(in, &a, &b, &c)
			for i := a; i < b+1; i++ {
				s[i] += c
			}
			break
		case 2:
			fmt.Fscan(in, &a, &b, &c)
			for i := a; i < b+1; i++ {
				t[i+n] += c
			}
			break
		case 3:
			fmt.Fscan(in, &a, &b, &c, &d)
			solve(a, b, c, d)
			fmt.Fprintln(out, r1, r2)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
