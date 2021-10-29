package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	c := make([]int, 3*n)
	for i := 0; i < n; i++ {
		j := i + 1
		if i == n-1 {
			j = 0
		}
		c[i] = b[i] ^ b[j]
	}
	for i := 0; i < n; i++ {
		j := i + 1
		if i == n-1 {
			j = 0
		}
		c[i+n] = a[i] ^ a[j]
		c[i+2*n] = a[i] ^ a[j]
	}

	d := ZAlgorithm(c)
	for i := n; i < 2*n; i++ {
		if d[i] >= n {
			fmt.Println(i-n, a[i-n]^b[0])
		}
	}
}

func ZAlgorithm(s []int) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		z[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for ; i+k < n && k+z[k] < j; k++ {
			z[i+k] = z[k]
		}
		i, j = i+k, j-k
	}
	return z
}
