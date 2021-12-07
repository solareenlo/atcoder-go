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

	p := make([]int, n)
	for i := range p {
		fmt.Fscan(in, &p[i])
		p[i]--
	}

	q := make([]int, n)
	for i := range q {
		q[p[i]] = i + 1
	}

	for i := range q {
		fmt.Fprint(out, q[i], " ")
	}
	fmt.Fprintln(out)
}
