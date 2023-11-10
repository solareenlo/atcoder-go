package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5555

	var a, b, l [N]int

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var n int
		fmt.Fscan(in, &n)
		t := 0
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &b[i])
		}
		for i := 1; i <= n; i++ {
			if b[i] != b[i-1] {
				t++
				l[t] = b[i]
			}
		}
		if l[t] == l[1] && t > 1 {
			t -= 1
		}
		var p int
		if t == n {
			p = 1
			for i := 1; i <= n; i++ {
				if a[i] == b[i] {
					p &= 1
				} else {
					p &= 0
				}
			}
		} else {
			p = 0
			for i := 1; i <= n; i++ {
				z := 1
				for j := 0; j < n; j++ {
					for z <= t && a[(i+j-1)%n+1] == l[z] {
						z++
					}
				}
				if z > t {
					p |= 1
				} else {
					p |= 0
				}
			}
		}
		if p != 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
