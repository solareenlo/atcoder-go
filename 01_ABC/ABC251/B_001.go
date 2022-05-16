package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n+1)
	b := make([]int, 3000006)
	for i := 1; i <= n; i = i + 1 {
		fmt.Fscan(in, &a[i])
		b[a[i]] = 1
	}

	for i := 0; i <= n; i = i + 1 {
		for j := i + 1; j <= n; j = j + 1 {
			for k := j + 1; k <= n; k = k + 1 {
				b[a[i]+a[j]+a[k]] = 1
			}
		}
	}

	ans := 0
	for i := 1; i <= m; i = i + 1 {
		ans = ans + b[i]
	}
	fmt.Println(ans)
}
