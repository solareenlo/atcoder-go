package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n int
	fmt.Fscan(in, &n)
	o, s := 0, 0
	var m, l, r [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &m[i], &l[i], &r[i])
		t := m[i] % (l[i] + r[i])
		s ^= max(0, t-1) / l[i]
		if t%l[i] == 0 || (m[i] < l[i]+r[i] && (t == l[i] || t >= 2*l[i])) {
			if o == 0 {
				o = i
			} else {
				fmt.Println("Yuri")
				return
			}
		}
	}
	if o == 0 {
		fmt.Println("Muri")
		return
	}
	if s == 0 {
		fmt.Println("Yuri")
		return
	}
	t := m[o] % (l[o] + r[o])
	s ^= max(0, t-1) / l[o]
	if s > (l[o]+r[o]-1)/l[o] {
		fmt.Println("Yuri")
		return
	}
	if (s+1)*l[o] < t {
		fmt.Println("Muri")
		return
	}
	if m[o] < l[o]+r[o] {
		fmt.Println("Yuri")
		return
	}
	if l[o]+t <= s*l[o] && s*l[o] < r[o]+t {
		fmt.Println("Muri")
	} else {
		fmt.Println("Yuri")
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
