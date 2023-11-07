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

	const SIZE = 100100

	pos := make([][]int, SIZE)

	var N int
	fmt.Fscan(in, &N)
	for n := 1; n <= 3*N; n = n + 1 {
		var v int
		fmt.Fscan(in, &v)
		pos[v] = append(pos[v], n)
	}
	ord := make([]int, 3*SIZE+1)
	for i := range ord {
		ord[i] = -1
	}
	for n := 1; n <= N; n = n + 1 {
		ord[pos[n][1]] = n
	}

	for n := 1; n <= 3*N; n = n + 1 {
		if ord[n] != -1 {
			fmt.Fprintf(out, "%d ", ord[n])
		}
	}
}
