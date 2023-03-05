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

	type pair struct {
		x byte
		y int
	}

	var N int
	fmt.Fscan(in, &N)

	p := make([]int, 405)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i])
	}

	ans := make([]pair, 0)

	var f func(byte, int)
	f = func(c byte, i int) {
		ans = append(ans, pair{c, i + 1})
		p[i], p[i+1+int(c-'A')] = p[i+1+int(c-'A')], p[i]
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N-2; j++ {
			if p[j]%2 != p[j+2]%2 && p[j]%2 != j%2 {
				f('B', j)
			}
		}
	}
	for i := 0; i < N-1; i++ {
		if p[i]%2 != p[i+1] && p[i]%2 == i%2 {
			f('A', i)
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N-2; j++ {
			if p[j] > p[j+2] {
				f('B', j)
			}
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%c %d\n", x.x, x.y)
	}
}
