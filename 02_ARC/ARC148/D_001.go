package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	n <<= 1
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	if (m & 1) == 0 {
		m >>= 1
		tmp := 0
		for i := 0; i < n; i++ {
			if a[i] >= m {
				a[i] -= m
				tmp++
			}
		}
		if (tmp & 1) != 0 {
			fmt.Println("Alice")
			return
		}
	}
	sort.Ints(a)
	for i := 0; i < n>>1; i++ {
		if a[i<<1]^a[i<<1|1] != 0 {
			fmt.Println("Alice")
			return
		}
	}
	fmt.Println("Bob")
}
