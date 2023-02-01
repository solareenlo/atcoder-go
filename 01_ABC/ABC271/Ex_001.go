package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var s string
var a [100]int
var x, y int

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		var tmp string
		fmt.Fscan(in, &x, &y, &tmp)
		s = " " + tmp
		Init()
		Solve()
		t--
	}
}

func Init() {
	for i := 1; i <= 8; i++ {
		if s[i] == '1' {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
	if x < 0 {
		a[1], a[5] = a[5], a[1]
		a[2], a[4] = a[4], a[2]
		a[6], a[8] = a[8], a[6]
		x *= -1
	}
	if y < 0 {
		a[2], a[8] = a[8], a[2]
		a[3], a[7] = a[7], a[3]
		a[4], a[6] = a[6], a[4]
		y *= -1
	}
	if x > y {
		a[1], a[3] = a[3], a[1]
		a[4], a[8] = a[8], a[4]
		a[5], a[7] = a[7], a[5]
		x, y = y, x
	}
}

func Solve() {
	res := int(1e18)
	if x == y && y == 0 {
		res = 0
	}
	if a[2] != 0 {
		if x == y {
			res = y
		}
		if a[3] != 0 {
			res = y
		}
		if a[5] != 0 {
			res = min(res, 2*y-x)
		}
		if a[4] != 0 {
			if ((x + y) & 1) == 0 {
				res = y
			}
			if a[5] != 0 {
				res = min(res, y+1)
			}
			if a[7] != 0 {
				res = min(res, y+2)
			}
			if a[1] != 0 {
				res = min(res, y+1)
			}
		}
	}
	if a[3] != 0 {
		if x == 0 {
			res = y
		}
		if a[1] != 0 {
			res = min(res, x+y)
		}
		if a[8] != 0 {
			res = min(res, 2*x+y)
		}
	}
	if a[4] != 0 && a[1] != 0 {
		res = min(res, 2*y+x)
	}
	if res == 1e18 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
