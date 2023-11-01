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

	const N = 2000010

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	a := make([]byte, n)
	b := make([]byte, n*2)
	for i := 0; i < n; i++ {
		a[i] = s[i]
		b[n+i] = s[n*2-i-1]
		b[i] = b[n+i]
	}

	x := strings.Index(string(b), string(a))
	if x == -1 {
		fmt.Println("-1")
	} else {
		i := n - x
		for j := 0; j < i; j++ {
			fmt.Fprintf(out, string(s[j]))
		}
		for j := n + i; j < 2*n; j++ {
			fmt.Fprintf(out, string(s[j]))
		}
		fmt.Fprintf(out, "\n%d", i)
	}
}
