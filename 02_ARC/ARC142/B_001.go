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

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprintf(out, "%d ", i*N+((N+1)/2)*(j%2)+j/2+1)
		}
		fmt.Fprintln(out)
	}
}
