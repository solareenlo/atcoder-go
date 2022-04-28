package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	E := make(map[int]int)
	O := make(map[int]int)
	for i := 0; i < n-1; i++ {
		if i%2 == 0 {
			E[a[i+1]-a[i]]++
		} else {
			O[a[i+1]-a[i]]++
		}
	}

	D := 0
	for j := 0; j < q; j++ {
		var Type int
		fmt.Fscan(in, &Type)
		if Type == 1 {
			var v int
			fmt.Fscan(in, &v)
			D += v
		} else if Type == 2 {
			var v int
			fmt.Fscan(in, &v)
			D -= v
		} else {
			var i, v int
			fmt.Fscan(in, &i, &v)
			i--
			if i > 0 {
				diff := a[i] - a[i-1]
				if i%2 == 1 {
					E[diff]--
					E[diff+v]++
				} else {
					O[diff]--
					O[diff+v]++
				}
			}
			if i < n-1 {
				diff := a[i+1] - a[i]
				if i%2 == 0 {
					E[diff]--
					E[diff-v]++
				} else {
					O[diff]--
					O[diff-v]++
				}
			}
			a[i] += v
		}
		fmt.Println(E[D] + O[-D])
	}
}
