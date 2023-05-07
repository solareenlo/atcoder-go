package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)

	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
	}

	ans := make([]pair, 0)
	for i := 1; i <= N; i++ {
		p := -1
		for j := 0; j < N; j++ {
			if P[j] == i {
				p = j
			}
		}
		for j := p - 1; j >= 0; j-- {
			if P[j] >= p+1 {
				ans = append(ans, pair{j + 1, p + 1})
				P[p], P[j] = P[j], P[p]
				p = j
			}
		}
	}

	fmt.Println(len(ans))
	for _, i := range ans {
		fmt.Println(i.x, i.y)
	}
}
