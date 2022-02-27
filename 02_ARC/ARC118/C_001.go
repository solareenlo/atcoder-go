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

	res := []int{6, 10, 15}
	for i := 16; len(res) < n; i++ {
		if i%6 == 0 || i%10 == 0 || i%15 == 0 {
			res = append(res, i)
		}
	}

	for _, x := range res {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
