package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 7

var c, th, fa [10]int

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}

	if c[1] == 0 {
		fmt.Println("NO")
		return
	}

	c[1]--

	th[1] = c[1] + c[2] - c[3] + c[4] - c[5] + c[6] - c[7]

	for i := 2; i <= n; i++ {
		th[i] = 2*c[i] - th[i-1]
	}

	for i := 1; i <= n; i++ {
		if th[i] < 0 {
			fmt.Println("NO")
			return
		}
	}

	for i := 1; i <= n; i++ {
		fa[i] = i
	}

	for i := 1; i < n; i++ {
		if th[i] != 0 {
			fa[get(i)] = get(i + 1)
		}
	}

	if th[n] != 0 {
		fa[get(n)] = get(1)
	}

	for i := 1; i <= n; i++ {
		if c[i] != 0 && get(i) != get(1) {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}

func get(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = get(fa[x])
	return fa[x]
}
