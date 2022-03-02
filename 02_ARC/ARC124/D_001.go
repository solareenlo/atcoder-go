package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n+m+1)
	z := 0
	for i := 1; i <= n+m; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == i {
			z++
		}
	}

	vis := make([]bool, n+m+1)
	x, y := 0, 0
	for i := 1; i <= n+m; i++ {
		if a[i] == i {
			continue
		}
		if vis[i] {
			continue
		}
		fn := 0
		fm := 0
		j := i
		for true {
			vis[j] = true
			if j <= n {
				fn = 1
			} else {
				fm = 1
			}
			if !vis[a[j]] {
				j = a[j]
			} else {
				break
			}
		}
		if fn == 0 {
			y++
		} else if fm == 0 {
			x++
		} else {
			z++
		}
	}

	fmt.Println(n + m - z + abs(y-x))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
