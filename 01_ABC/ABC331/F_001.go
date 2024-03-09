package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 1000005
const P = 131

func lowbit(x int) int {
	return x & (-x)
}

var n int
var dat [2][N]int

func change(op, x, v int) {
	for ; x <= n; x += lowbit(x) {
		dat[op][x] += v
	}
	return
}

func query(op, x int) int {
	res := 0
	for ; x > 0; x -= lowbit(x) {
		res += dat[op][x]
	}
	return res
}

func ask(op, l, r int) int { return query(op, r) - query(op, l-1) }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &n, &m)
	var p [N]int
	p[0] = 1
	for i := 1; i <= 1000000; i++ {
		p[i] = p[i-1] * P
	}
	var tmp string
	fmt.Fscan(in, &tmp)
	tmp = " " + tmp
	s := strings.Split(tmp, "")
	for i := 1; i <= n; i++ {
		x := int(s[i][0])
		y := int(s[n-i+1][0])
		change(0, i, x*p[i-1])
		change(1, i, y*p[i-1])
	}
	for m > 0 {
		m--
		var z int
		fmt.Fscan(in, &z)
		if (z & 1) != 0 {
			var x int
			fmt.Fscan(in, &x)
			var c string
			fmt.Fscan(in, &c)
			y := int(c[0])
			change(0, x, (y-int(s[x][0]))*p[x-1])
			change(1, n-x+1, (y-int(s[x][0]))*p[n-x])
			s[x] = c
		} else {
			var x, y int
			fmt.Fscan(in, &x, &y)
			if x == y {
				fmt.Fprintln(out, "Yes")
				continue
			}
			z = (y - x - 1) >> 1
			hsh1 := ask(0, x, z+x)
			hsh2 := ask(1, n-y+1, z+n-y+1)
			if x-1 > n-y {
				hsh2 *= p[x-1-(n-y)]
			} else {
				hsh1 *= p[(n-y)-(x-1)]
			}
			if hsh1^hsh2 != 0 {
				fmt.Fprintln(out, "No")
			} else {
				fmt.Fprintln(out, "Yes")
			}
		}
	}
}
