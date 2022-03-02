package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := n; i > 1; i-- {
		for j := i; j >= 1; j-- {
			now := 1
			a[i], a[j] = a[j], a[i]
			for k := 1; k < i; k++ {
				now = lcm(now, gcd(a[k], a[i]))
			}
			if now != a[i] {
				break
			} else if j == 1 {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
	fmt.Fprintln(out)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}
