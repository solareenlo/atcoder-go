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
	for i := n; i >= 0; i-- {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n+1)
	zs := make([]int, 0)
	for i := 2; i <= n; i++ {
		if f[i] == 0 {
			zs = append(zs, i)
			for j := i + i; j <= n; j += i {
				f[j] = 1
			}
		}
	}

	sum := abs(a[n])
	for i := 2; i*i <= sum; i++ {
		if sum%i == 0 {
			if i > n {
				zs = append(zs, i)
			}
			for sum%i == 0 {
				sum /= i
			}
		}
	}

	if sum > max(1, n) {
		zs = append(zs, sum)
	}

	for _, x := range zs {
		b := make([]int, n+1)
		copy(b, a)
		for j := n; j >= x; j-- {
			b[j-(x-1)] += b[j]
		}
		fl := 1
		for j := 0; j <= min(n, x-1) && fl != 0; j++ {
			tmp := 0
			if (b[j] % x) == 0 {
				tmp = 1
			}
			fl &= tmp
		}
		if fl != 0 {
			fmt.Fprintln(out, x)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
