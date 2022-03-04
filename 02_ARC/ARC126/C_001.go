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

	a := make([]int, n+1)
	c := make([]int, 1000005)
	mx := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
		c[a[i]]++
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += mx - a[i]
	}
	if sum <= k {
		fmt.Println((k-sum)/n + mx)
		return
	}

	s := make([]int, 1000005)
	for i := 1; i <= mx+mx; i++ {
		s[i] = s[i-1] + c[i]*i
		c[i] += c[i-1]
	}
	for i := mx; i > 0; i-- {
		sum := 0
		for j := i; j <= mx+i; j += i {
			sum += (c[j]-c[j-i])*j - (s[j] - s[j-i])
		}
		if sum <= k {
			fmt.Println(i)
			return
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
