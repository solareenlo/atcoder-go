package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	s := 0
	var a [200000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		s += a[i]
	}
	t := 0
	if s == x+y {
		for l, r := 0, 0; l < n; l++ {
			for ; t < x; r++ {
				t += a[r%n]
			}
			if t == x {
				fmt.Println("Yes")
				return
			}
			t -= a[l]
		}
	}
	fmt.Println("No")
}
