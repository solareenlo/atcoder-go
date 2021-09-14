package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	i, j int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, n, a, b int
	fmt.Fscan(in, &h, &w, &n)

	m := map[pair]int{}
	res := [11]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if 3 <= a+j && a+j <= h && 3 <= b+k && b+k <= w {
					res[m[pair{a + j, b + k}]]++
					m[pair{a + j, b + k}]++
				}
			}
		}
	}

	fmt.Println((h-2)*(w-2) - res[0])
	for i := 0; i < 9; i++ {
		fmt.Println(res[i] - res[i+1])
	}
}
