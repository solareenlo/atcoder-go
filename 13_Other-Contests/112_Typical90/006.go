package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K int
	var S string
	fmt.Fscan(in, &N, &K, &S)

	r := N - K
	var P []string
	for _, c := range S {
		for len(P) != 0 && P[len(P)-1] > string(c) && r != 0 {
			r--
			P = P[:len(P)-1]
		}
		P = append(P, string(c))
	}
	for r != 0 {
		r--
		P = P[:len(P)-1]
	}
	fmt.Fprintln(out, strings.Join(P, ""))
}
