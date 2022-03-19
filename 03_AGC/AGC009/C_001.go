package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
	}

	if a > b {
		a, b = b, a
	}

	for i := 2; i < n; i++ {
		if s[i+1]-s[i-1] < a {
			fmt.Println(0)
			return
		}
	}

	f := make([]int, n+1)
	const mod = 1_000_000_007
	f[0] = 1
	s[0] = -b
	sum := 0
	r := 0
	for i := 1; i <= n; i++ {
		for ; r < i && s[i]-s[r] >= b; r++ {
			sum = (sum + f[r]) % mod
		}
		f[i] = sum
		if s[i]-s[i-1] < a {
			sum = 0
			r = i - 1
		}
	}

	for i := r; i <= n; i++ {
		sum = (sum + f[i]) % mod
	}
	fmt.Println(sum)
}
