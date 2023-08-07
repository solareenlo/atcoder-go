package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	num := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[b] = append(g[b], a)
		num[a]++
	}
	res := make([]float64, n)
	res2 := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = 1.0
		res2[i] = 1.0
	}
	for i := 0; i < 500; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for _, a := range g[j] {
				sum += res2[a] / float64(num[a])
			}
			res[j] = 0.1 + 0.9*sum
		}
		res, res2 = res2, res
	}
	for _, a := range res {
		fmt.Println(a)
	}
}
