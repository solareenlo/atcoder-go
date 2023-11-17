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

	for i := 1; i <= 365; i++ {
		for j := 1; j <= 26; j++ {
			var t int
			fmt.Fscan(in, &t)
		}
		fmt.Fprintln(out, (i%26)+1)
		out.Flush()
	}
}
