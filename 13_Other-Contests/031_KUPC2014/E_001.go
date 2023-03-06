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

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var n, m int
		fmt.Fscan(in, &n, &m)
		if check(n, m) {
			fmt.Fprintln(out, "Possible")
		} else {
			fmt.Fprintln(out, "Impossible")
		}
	}
}

func check(n, m int) bool {
	if n*m%8 != 0 || n == 1 || m == 1 {
		return false
	}
	if (n == 2 && m == 4) || (n == 4 && m == 2) {
		return false
	}
	return true
}
