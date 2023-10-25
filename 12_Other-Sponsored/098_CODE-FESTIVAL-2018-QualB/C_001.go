package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	var tmp [5]int = [5]int{1, 3, 0, 2, 4}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (tmp[i%5] == j%5 || (i == 0 && j%5 == 4) || (j == 0 && i%5 == 4)) || (i == n-1 && j%2 == 0) || (j == n-1 && i%2 == 0) {
				fmt.Fprint(out, "X")
			} else {
				fmt.Fprint(out, ".")
			}
		}
		fmt.Fprintln(out)
	}
}
