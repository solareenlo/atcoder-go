package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const INF = 1e18

func main() {
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		solve()
		T--
	}
}

func solve() {
	var p, a, b, s, g int
	fmt.Fscan(in, &p, &a, &b, &s, &g)
	mp := make(map[int]int)
	if s == g {
		fmt.Fprintln(out, 0)
		return
	}
	if a == 0 {
		if b == g {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, -1)
		}
		return
	}
	t := int(math.Sqrt(float64(p)))
	c, d := 1, 0
	for i := 1; i <= t; i++ {
		c = c * a % p
		d = (d*a + b) % p
	}
	val, res := s, int(INF)
	for i := t; i-t <= p; i += t {
		val = (val*c + d) % p
		if _, ok := mp[val]; ok == false {
			mp[val] = i
		}
	}
	val = g
	for i := 1; i <= t; i++ {
		val = (val*a + b) % p
		if _, ok := mp[val]; ok {
			res = min(res, mp[val]-i)
		}
	}
	if res == INF {
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
