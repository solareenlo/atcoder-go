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

	var n int
	fmt.Fscan(in, &n)
	for n > 0 {
		n--
		var x int
		fmt.Fscan(in, &x)
		if x%2 == 0 {
			fmt.Fprintf(out, "%d ", x)
		}
	}
}
