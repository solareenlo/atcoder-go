package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscan(in, &n, &l)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}

	ans := 1
	x := 0
	R, L := 0, 0
	for i := 1; i <= n; i++ {
		var t int
		fmt.Fscan(in, &t)
		ans += t / (2 * l)
		t %= 2 * l
		if t == 0 {
			continue
		}
		ans++
		u := t <= (l-p[i])<<1
		tmp := 0
		if u {
			tmp = 1
		}
		v := t <= p[i]<<1
		if i == n {
			ans -= tmp
		} else if v {
			L++
			if u {
				x++
			}
		} else if u {
			R++
		}
		for L-x < R+x {
			if x != 0 {
				x--
			} else {
				R--
			}
		}
	}

	fmt.Println((ans - R - x) * l * 2)
}
