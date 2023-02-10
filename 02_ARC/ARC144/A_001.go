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

	fmt.Println(2 * n)
	if n%4 != 0 {
		fmt.Fprint(out, n%4)
	}
	n /= 4
	for n > 0 {
		n--
		fmt.Fprint(out, 4)
	}
}
