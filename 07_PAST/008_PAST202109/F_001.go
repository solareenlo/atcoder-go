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

	var N int
	var S string
	fmt.Fscan(in, &N, &S)

	ids := make([]int, 0)
	for i := 0; i < N; i++ {
		if S[i] == '0' {
			ids = append(ids, i)
		}
	}

	if len(ids) == 1 {
		fmt.Fprintln(out, -1)
	} else {
		sz := 0
		for i := 0; i < N; i++ {
			if S[i] == '1' {
				fmt.Fprint(out, i+1)
			} else {
				sz++
				fmt.Fprint(out, ids[sz%len(ids)]+1)
			}
			if i+1 == N {
				fmt.Fprintln(out)
			} else {
				fmt.Fprint(out, " ")
			}
		}
	}
}
