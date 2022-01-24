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

	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}

	a[0] = n + 1
	cnt := 1
	b := make([]byte, n+1)
	for i := 0; i < 26; i++ {
		for j := 0; j <= n && cnt <= n; j++ {
			if a[cnt]+1 == a[j] {
				b[a[cnt]] = 'A' + byte(i)
				cnt++
			}
		}
	}

	if cnt <= n {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 1; i < n+1; i++ {
		fmt.Fprint(out, string(b[i]))
	}
	fmt.Fprintln(out)
}
