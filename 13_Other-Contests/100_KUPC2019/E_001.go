package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 100005

	var n int
	fmt.Fscan(in, &n)

	var p [MX]int
	a := 0
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		d := make([]int, m+1)
		for j := range d {
			d[j] = -1
		}
		for j := 0; j < m-1; j++ {
			fmt.Fscan(in, &p[j+2])
		}
		for j := m; j > 0; j-- {
			if d[j] == -1 {
				d[j] = 0
			}
			if d[p[j]] != 3 {
				if d[p[j]] == d[j] || d[j] == 2 {
					d[p[j]] = 3
				} else if d[p[j]]&2 != 0 {
					d[p[j]] = d[j] ^ 1
				}
			}
		}
		if d[1] == 2 {
			a = 2
		}
		if d[1] == 1 {
			a ^= 1
		} else {
			a ^= 0
		}
	}

	if a != 0 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
