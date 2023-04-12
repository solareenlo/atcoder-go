package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(in, &m, &n)
	var C [1 << 17]int
	var D [1 << 17]int
	for i := 0; i < m; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < n; j++ {
			C[j] += int(c[j] - '0')
			D[i] += int(c[j] - '0')
		}
	}

	ans := 1
	for j := 0; j < 2; j++ {
		S := make([]int, 2*n)
		T := make([]int, 2*n)
		for i := 0; i < n; i++ {
			T[n+i] = C[i]
			T[n-1-i] = C[i]
			S[2*n-1-i] = C[i]
			S[i] = C[i]
		}
		X := ZAlgorithm(S)
		Y := ZAlgorithm(T)
		k := 0
		for i := 0; i < n-1; i++ {
			if X[2*n-1-i] == i+1 && Y[n+1+i] == n-1-i {
				k++
			}
		}
		ans *= k
		m, n = n, m
		C, D = D, C
	}
	fmt.Println(ans)
}

func ZAlgorithm(s []int) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		z[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for ; i+k < n && k+z[k] < j; k++ {
			z[i+k] = z[k]
		}
		i, j = i+k, j-k
	}
	return z
}
