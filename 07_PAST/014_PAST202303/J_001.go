package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S, T [1 << 19]string

	var h, w int
	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &S[i])
	}
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &T[i])
	}
	V := make([][]string, 3*w)
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			V[j] = append(V[j], string(S[i][j]))
			V[j+w] = append(V[j+w], string(T[i][j]))
			V[j+w+w] = append(V[j+w+w], string(T[i][j]))
		}
	}
	v := make([]string, 3*w)
	for i := range v {
		v[i] = strings.Join(V[i], "")
	}
	Z := z_algorithm(v)
	for j := w; j < 2*w; j++ {
		if Z[j] >= w {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func z_algorithm(s []string) []int {
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
