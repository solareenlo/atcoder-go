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
		fmt.Fprintln(out, "Yes")
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i == j {
					if j == N-1 {
						fmt.Fprintln(out, (N-i-1)*(N+1)+1)
					} else {
						fmt.Fprintf(out, "%d ", (N-i-1)*(N+1)+1)
					}
				} else {
					if j == N-1 {
						fmt.Fprintln(out, i*N+j+1)
					} else {
						fmt.Fprintf(out, "%d ", i*N+j+1)
					}
				}
			}
		}
	} else {
		fmt.Fprintln(out, "No")
	}
}
