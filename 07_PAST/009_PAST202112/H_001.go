package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S, T string
	fmt.Fscan(in, &S, &T)

	d := make([]int, 5050)
	for i := 0; i < len(S); i++ {
		p := 0
		for j := 0; j < len(T); j++ {
			q := max(p, d[j])
			if S[i] != T[j] {
				p++
			}
			d[j] = max(p, d[j])
			p = q
		}
	}

	ans := 0
	for j := 0; j < len(T); j++ {
		ans = max(ans, d[j])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
