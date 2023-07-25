package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S, T, U string
	fmt.Fscan(in, &S, &T, &U)

	var a, b, c [26]int
	N := len(S)
	for i := 0; i < N; i++ {
		a[(int)(S[i]-'A')]++
		b[(int)(T[i]-'A')]++
		c[(int)(U[i]-'A')]++
	}
	L, R := 0, 0
	for i := 0; i < 26; i++ {
		if c[i] > b[i] {
			L += c[i] - b[i]
		}
		R += min(a[i], c[i])
	}
	if L <= N/2 && N/2 <= R {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
