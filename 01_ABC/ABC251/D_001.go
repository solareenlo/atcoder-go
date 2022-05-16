package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintln(out, 300)
	for i := 1; i <= 100; i++ {
		fmt.Fprintf(out, "%d %d00 %d0000 ", i, i, i)
	}
}
