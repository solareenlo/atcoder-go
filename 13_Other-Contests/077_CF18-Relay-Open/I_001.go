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
	var ct [5]int
	s := 0
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		s += p
		ct[p%5]++
	}
	m := min(ct[3], ct[4])
	ct[3] -= m
	ct[4] -= m
	x := m + ct[2]
	y := ct[1]
	z := 0
	w := 0
	y += ct[3] / 2
	w = ct[3] % 2
	x += (ct[4] / 3)
	z += (ct[4] % 3)
	x += (y / 2)
	y %= 2
	s -= (2*x + y - z - 2*w)
	ans := x
	if z == 0 && w == 0 {
		ans += y
	} else {
		c := -z - 2*w + y
		if 6+c <= 2 {
			ans -= 2
		} else if 4+c <= 2 {
			ans--
		}
	}
	if ans <= 0 {
		ans = 1
	}
	fmt.Println(s, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
