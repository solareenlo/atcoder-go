package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n+3)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := n + 2; i > 0; i-- {
		a[i] -= a[i-1]
	}
	for i := n + 2; i > 0; i-- {
		a[i] -= a[i-1]
	}
	for i := n + 2; i > 0; i-- {
		a[i] -= a[i-1]
	}

	for i := 0; i < n+3; i++ {
		if a[i] < 0 {
			fmt.Println("No")
			return
		}

		d := a[i]
		if i+2*k+1 < n+3 {
			a[i] -= d
			a[i+1] -= d
			a[i+k] += 4 * k * d
			a[i+k+1] -= 4 * k * d
			a[i+2*k] += d
			a[i+2*k+1] += d
		} else {
			if d != 0 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
