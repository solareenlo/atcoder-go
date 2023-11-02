package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)

	p := []int{4, 9, 5, 7, 11, 13, 17, 19, 23}

	n := 108
	fmt.Fprintln(out, n)
	out.Flush()
	res, t, w := 0, 0, 0
	for _, i := range p {
		for j := t + 1; j < t+i; j++ {
			fmt.Fprintf(out, "%d ", j+1)
			out.Flush()
		}
		fmt.Fprintf(out, "%d ", t+1)
		out.Flush()
		t += i
	}
	fmt.Fprintln(out)
	out.Flush()
	t = 1
	var b [333]int
	for i, k := 1, 0; i <= n; i++ {
		fmt.Scan(&b[i])
		if i == k+1 {
			for res%p[w] != b[i]-i {
				res += t
			}
			k += p[w]
			t *= p[w]
			w++
		}
	}
	fmt.Fprintln(out, res)
	out.Flush()
}
