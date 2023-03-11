package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Scan(&N)

	if N%2 != 0 {
		fmt.Fprintln(out, (N/2)*N)
		for d := 1; d <= N/2; d++ {
			for i := 0; i < N; i++ {
				fmt.Fprintln(out, i+1, (i+d)%N+1, (i+d*2)%N+1)
			}
		}
	} else {
		fmt.Fprintln(out, N*(N-1))
		N--
		for i := 0; i < N; i++ {
			fmt.Fprintln(out, N+1, i+1, (i+1)%N+1)
			fmt.Fprintln(out, N+1, i+1, (i+N/2)%N+1)
			fmt.Fprintln(out, N+1, i+1, (i+N/2)%N+1)
		}
		for d := 1; d <= N/2; d++ {
			for i := 0; i < N; i++ {
				fmt.Fprintln(out, i+1, (i+d)%N+1, (i+d*2)%N+1)
				if d < N/2 {
					fmt.Fprintln(out, i+1, (i+d)%N+1, (i+d*2)%N+1)
				}
			}
		}
	}
}
