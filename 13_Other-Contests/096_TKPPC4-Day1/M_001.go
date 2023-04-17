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

	var N int
	fmt.Fscan(in, &N)
	la, ra := 0, N
	for ra-la > 1 {
		m := (la + ra) >> 1
		if N >= 4096 && la == 0 && ra == N {
			m = 4096
		}
		fmt.Fprintf(out, "? %d", m)
		out.Flush()
		for i := 0; i < m; i++ {
			fmt.Fprintf(out, " %d", i+1)
			out.Flush()
		}
		fmt.Fprintln(out)
		out.Flush()
		var s string
		fmt.Fscan(in, &s)
		if s == "Rabbit" {
			ra = m
		} else {
			la = m
		}
	}
	lb, rb := 0, ra
	for rb-lb > 1 {
		m := (lb + rb) >> 1
		fmt.Fprintf(out, "? %d", ra-m)
		out.Flush()
		for i := m; i < ra; i++ {
			fmt.Fprintf(out, " %d", i+1)
			out.Flush()
		}
		fmt.Fprintln(out)
		out.Flush()
		var s string
		fmt.Fscan(in, &s)
		if s == "Rabbit" {
			lb = m
		} else {
			rb = m
		}
	}
	fmt.Fprintln(out, "!", lb+1, la+1)
	out.Flush()
}
