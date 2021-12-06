package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	var l int
	r := make([]int, M+1)
	y := make([]int, M+1)
	n := make([]int, 200001)
	m := make([]int, 200001)
	b := make([]int, 200001)
	for i := 1; i < M+1; i++ {
		fmt.Fscan(in, &l, &r[i], &y[i])
		y[i] = r[i] - y[i] - l + 1
		if b[l] != 0 {
			n[b[l]] = i
		} else {
			m[l] = i
		}
		b[l] = i
	}

	k := 0
	c := make([]int, 200001)
	for i := 1; i < N+1; i++ {
		for j := m[i]; j > 0; j = n[j] {
			if r[j] > c[y[j]+k] {
				c[y[j]+k] = r[j]
			}
		}
		if c[k] >= i {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
			k++
		}
	}
	fmt.Println()
}
