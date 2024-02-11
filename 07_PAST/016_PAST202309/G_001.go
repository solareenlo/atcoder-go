package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var a [12]int

func f(t int) {
	if t == (1<<(n*3))-1 {
		m++
		return
	}

	j := 0

	for ((1 << j) & t) != 0 {
		j++
	}

	for k := j + 1; k < n*3-1; k++ {
		if (1<<k)&t == 0 {
			for l := k + 1; l < n*3; l++ {
				if (1<<l)&t == 0 {
					if abs(a[j]-a[k]) < a[l] && a[l] < a[j]+a[k] {
						f(t + (1 << j) + (1 << k) + (1 << l))
					}
				}
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 0; i < n*3; i++ {
		fmt.Fscan(in, &a[i])
	}

	f(0)
	fmt.Println(m)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
