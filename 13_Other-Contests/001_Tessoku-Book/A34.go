package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100001

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	var G [N]int
	for i := 0; i < N; i++ {
		var T [3]bool
		for j := range T {
			T[j] = true
		}
		if i >= x {
			T[G[i-x]] = false
		}
		if i >= y {
			T[G[i-y]] = false
		}
		for j := 0; j < 3; j++ {
			if T[j] {
				G[i] = j
				break
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans ^= G[a]
	}
	if ans != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
