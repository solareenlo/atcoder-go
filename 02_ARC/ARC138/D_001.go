package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Scan(&n, &k)

	if k%2 == 0 {
		fmt.Fprintln(out, "No")
		return
	}
	if 1 < k && n == k {
		fmt.Fprintln(out, "No")
		return
	}

	z := make([]int, 0)
	for i := 0; i < k-1; i++ {
		z = append(z, mask(k+1)^(1<<i))
	}
	for i := k - 1; i < n; i++ {
		z = append(z, mask(k-1)^(1<<i))
	}

	s := 1 << n
	a := make([]int, s)
	for i := 0; i < s; i++ {
		v := i ^ (i >> 1)
		for j := 0; j < n; j++ {
			if v&(1<<j) != 0 {
				a[i] ^= z[j]
			}
		}
	}
	fmt.Fprintln(out, "Yes")
	for i := 0; i < len(a); i++ {
		fmt.Fprint(out, a[i], " ")
	}
	fmt.Fprintln(out)
}

func mask(i int) int {
	return (1 << i) - 1
}
