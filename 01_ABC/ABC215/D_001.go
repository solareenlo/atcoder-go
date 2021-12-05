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

	var n, m int
	fmt.Fscan(in, &n, &m)

	const SIZE = 100005
	v := make([]int, SIZE)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		v[a] = 1
	}

	for i := 1; i < SIZE; i++ {
		for j := i; j < SIZE; j += i {
			v[i] |= v[j]
		}
	}

	v[1] = 0
	for i := 1; i < SIZE; i++ {
		for j := i; j < SIZE; j += i {
			v[j] |= v[i]
		}
	}

	z := 0
	for i := 1; i < m+1; i++ {
		if v[i] == 0 {
			z++
		}
	}

	fmt.Fprintln(out, z)
	for i := 1; i < m+1; i++ {
		if v[i] == 0 {
			fmt.Fprintln(out, i)
		}
	}
}
