package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func div(M, N *string) {
	m := strings.Split(*M, "")
	n := strings.Split(*N, "")
	i := 0
	for ; i < len(m); i++ {
		if ^m[i][0]&1 != 0 {
			continue
		}
		if i+len(n) > len(m) {
			break
		}
		for j := 1; j < len(n); j++ {
			m[i+j] = string(m[i+j][0] ^ n[j][0]&1)
		}
	}
	*M = strings.Join(m, "")[i:]
}

func gcd(m, n *string) {
	for len(*n) > 0 {
		div(m, n)
		*m, *n = *n, *m
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	var X string
	fmt.Fscan(in, &N, &X)

	var A string
	for i := 0; i < N; i++ {
		var B string
		fmt.Fscan(in, &B)
		gcd(&A, &B)
	}

	mod := 998244353
	a := 0
	for i := 0; i+len(A) <= len(X); i++ {
		tmp := 0
		if i != len(X) && X[i]&1 != 0 {
			tmp = 1
		}
		a = (a*2 | tmp) % mod
	}

	Y := X
	div(&Y, &A)

	tmp := 0
	if len(Y) == 0 || (len(X) >= len(Y) && X[len(X)-len(Y)]&1 != 0) {
		tmp = 1
	}
	fmt.Println((a + tmp) % mod)
}
